package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Service string `json:"service"`
	Message string `json:"message"`
	Website string `json:"website"` // honeypot: los humanos no lo ven ni lo llenan
}

type StoredMessage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Company   string    `json:"company,omitempty"`
	Service   string    `json:"service,omitempty"`
	Message   string    `json:"message"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"createdAt"`
}

func newContactHandler(store *MessageStore) http.Handler {
	limiter := newRateLimiter(5, 10*time.Minute)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		if !limiter.allow(ip) {
			writeJSON(w, http.StatusTooManyRequests, map[string]string{
				"error": "Has enviado demasiados mensajes. Intenta de nuevo en unos minutos.",
			})
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
		var req ContactRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Solicitud inválida."})
			return
		}

		success := map[string]any{
			"ok":      true,
			"message": "¡Gracias! Recibimos tu mensaje y te contactaremos muy pronto.",
		}

		if req.Website != "" {
			writeJSON(w, http.StatusCreated, success)
			return
		}

		req.normalize()
		if fields := req.validate(); len(fields) > 0 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
				"error":  "Revisa los campos marcados.",
				"fields": fields,
			})
			return
		}

		msg := StoredMessage{
			ID:        newID(),
			Name:      req.Name,
			Email:     req.Email,
			Phone:     req.Phone,
			Company:   req.Company,
			Service:   req.Service,
			Message:   req.Message,
			IP:        ip,
			CreatedAt: time.Now().UTC(),
		}
		if err := store.Save(msg); err != nil {
			log.Printf("error al guardar mensaje: %v", err)
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "No pudimos guardar tu mensaje. Intenta más tarde.",
			})
			return
		}

		log.Printf("nuevo mensaje de contacto %s (%s)", msg.ID, msg.Email)
		writeJSON(w, http.StatusCreated, success)
	})
}

func (c *ContactRequest) normalize() {
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(c.Email)
	c.Phone = strings.TrimSpace(c.Phone)
	c.Company = strings.TrimSpace(c.Company)
	c.Service = strings.TrimSpace(c.Service)
	c.Message = strings.TrimSpace(c.Message)
}

func (c *ContactRequest) validate() map[string]string {
	fields := map[string]string{}

	switch n := utf8.RuneCountInString(c.Name); {
	case n == 0:
		fields["name"] = "Ingresa tu nombre."
	case n > 100:
		fields["name"] = "El nombre es demasiado largo."
	}

	if c.Email == "" {
		fields["email"] = "Ingresa tu correo electrónico."
	} else if addr, err := mail.ParseAddress(c.Email); err != nil || addr.Address != c.Email || len(c.Email) > 150 {
		fields["email"] = "Ingresa un correo electrónico válido."
	}

	if utf8.RuneCountInString(c.Phone) > 30 {
		fields["phone"] = "El teléfono es demasiado largo."
	}
	if utf8.RuneCountInString(c.Company) > 120 {
		fields["company"] = "El nombre de la empresa es demasiado largo."
	}
	if c.Service != "" && !isKnownService(c.Service) {
		fields["service"] = "Selecciona un servicio válido."
	}

	switch n := utf8.RuneCountInString(c.Message); {
	case n < 10:
		fields["message"] = "Cuéntanos un poco más (mínimo 10 caracteres)."
	case n > 3000:
		fields["message"] = "El mensaje no puede superar 3000 caracteres."
	}

	return fields
}

func isKnownService(id string) bool {
	if id == "otro" {
		return true
	}
	for _, s := range siteContent.Services {
		if s.ID == id {
			return true
		}
	}
	return false
}

// MessageStore guarda los mensajes de contacto en un archivo JSON.
type MessageStore struct {
	mu   sync.Mutex
	path string
}

func NewMessageStore(path string) (*MessageStore, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return &MessageStore{path: path}, nil
}

func (s *MessageStore) Save(m StoredMessage) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var msgs []StoredMessage
	data, err := os.ReadFile(s.path)
	switch {
	case err == nil && len(data) > 0:
		if err := json.Unmarshal(data, &msgs); err != nil {
			return fmt.Errorf("leer %s: %w", s.path, err)
		}
	case err != nil && !errors.Is(err, os.ErrNotExist):
		return err
	}

	msgs = append(msgs, m)
	out, err := json.MarshalIndent(msgs, "", "  ")
	if err != nil {
		return err
	}

	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, out, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

type rateLimiter struct {
	mu     sync.Mutex
	hits   map[string][]time.Time
	limit  int
	window time.Duration
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{hits: map[string][]time.Time{}, limit: limit, window: window}
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)
	recent := rl.hits[key][:0]
	for _, t := range rl.hits[key] {
		if t.After(cutoff) {
			recent = append(recent, t)
		}
	}
	if len(recent) >= rl.limit {
		rl.hits[key] = recent
		return false
	}
	rl.hits[key] = append(recent, now)
	return true
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func newID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}
