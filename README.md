# NGN NetVision Technologies — Sitio web empresarial

Página única (one-page) corporativa con **React + Vite** en el frontend y **Go** (solo librería estándar) en el backend.

Secciones: Inicio (portada con fotos) · Servicios (Desarrollo de Software, Redes y Telecomunicaciones, Rastreo Vehicular GPS) · Nosotros · Proceso · Contacto.

## Estructura

```
backend/
  main.go      Servidor HTTP, rutas, CORS, sirve el build de React
  content.go   TODO el contenido del sitio (textos, servicios, cifras, contacto)
  contact.go   Formulario de contacto: validación, anti-spam y guardado
frontend/
  src/App.jsx          Carga el contenido desde la API y arma la página
  src/components/      Navbar, Hero, Stats, Services, About, Process, Contact, Footer
  src/index.css        Estilos y diseño responsivo
```

## Requisitos

- Go 1.22 o superior — https://go.dev/dl/
- Node.js 18 o superior

## Desarrollo

En dos terminales:

```bash
# Terminal 1 — API en http://localhost:8080
cd backend
go run .

# Terminal 2 — React en http://localhost:5173 (redirige /api al backend)
cd frontend
npm install
npm run dev
```

Abre http://localhost:5173

## Producción (un solo servidor)

```bash
cd frontend && npm run build      # genera frontend/dist
cd ../backend && go build -o nepkvision.exe .
./nepkvision.exe                  # sitio + API en http://localhost:8080
```

## API

| Método | Ruta           | Descripción                                     |
|--------|----------------|-------------------------------------------------|
| GET    | `/api/health`  | Estado del servidor                             |
| GET    | `/api/content` | Contenido completo del sitio (JSON)             |
| POST   | `/api/contact` | Recibe `{name, email, phone, company, service, message}` |

Los mensajes de contacto se guardan en `backend/data/messages.json`.
Límite: 5 mensajes por IP cada 10 minutos. Incluye un campo trampa (honeypot) contra bots.

## Variables de entorno (backend)

| Variable         | Por defecto              | Uso                                      |
|------------------|--------------------------|------------------------------------------|
| `ADDR`           | `:8080`                  | Puerto del servidor                      |
| `DATA_DIR`       | `data`                   | Carpeta donde se guardan los mensajes    |
| `STATIC_DIR`     | `../frontend/dist`       | Build de React a servir                  |
| `ALLOWED_ORIGIN` | `http://localhost:5173`  | Origen permitido por CORS                |

## Personalizar contenido

Edita `backend/content.go` y reinicia el servidor. **Antes de publicar**, reemplaza los datos de ejemplo:
correo, teléfono, dirección y las cifras (proyectos, clientes, vehículos).
