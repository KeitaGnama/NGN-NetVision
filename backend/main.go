package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	addr := envOr("ADDR", ":8080")
	dataDir := envOr("DATA_DIR", "data")
	staticDir := envOr("STATIC_DIR", filepath.Join("..", "frontend", "dist"))
	allowedOrigin := envOr("ALLOWED_ORIGIN", "http://localhost:5173")

	store, err := NewMessageStore(filepath.Join(dataDir, "messages.json"))
	if err != nil {
		log.Fatalf("no se pudo inicializar el almacenamiento: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", handleHealth)
	mux.HandleFunc("GET /api/content", handleContent)
	mux.Handle("POST /api/contact", newContactHandler(store))
	mux.Handle("/", spaHandler(staticDir))

	srv := &http.Server{
		Addr:              addr,
		Handler:           logRequests(cors(allowedOrigin, mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("NGN NetVision Technologies API escuchando en %s", addr)
	log.Fatal(srv.ListenAndServe())
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleContent(w http.ResponseWriter, r *http.Request) {
	// no-cache: el navegador siempre revisa si hay contenido nuevo, así los
	// cambios en content.go se ven en cuanto se reinicia el servidor.
	w.Header().Set("Cache-Control", "no-cache")
	writeJSON(w, http.StatusOK, siteContent)
}

// spaHandler sirve el build de React y devuelve index.html para cualquier ruta desconocida.
func spaHandler(dir string) http.Handler {
	files := http.FileServer(http.Dir(dir))
	index := filepath.Join(dir, "index.html")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Ruta no encontrada."})
			return
		}
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}
		if _, err := os.Stat(index); err != nil {
			http.Error(w, "Frontend no compilado: ejecuta 'npm run build' en /frontend o usa 'npm run dev'.", http.StatusNotFound)
			return
		}

		p := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+r.URL.Path)))
		if info, err := os.Stat(p); err == nil && !info.IsDir() {
			files.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, index)
	})
}

func cors(allowedOrigin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && (allowedOrigin == "*" || origin == allowedOrigin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Add("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("%s %s %d %s", r.Method, r.URL.Path, rec.status, time.Since(start).Round(time.Millisecond))
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("error al escribir JSON: %v", err)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
