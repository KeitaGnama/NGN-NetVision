// En desarrollo Vite redirige /api al backend de Go (ver vite.config.js).
// Si el frontend se publica en otro dominio, define VITE_API_URL.
const API_BASE = import.meta.env.VITE_API_URL ?? '';

export async function getContent() {
  // no-store: siempre pide la versión más reciente del contenido.
  const res = await fetch(`${API_BASE}/api/content`, { cache: 'no-store' });
  if (!res.ok) throw new Error(`Error ${res.status} al cargar el contenido`);
  return normalizeContent(await res.json());
}

// Asegura que los datos tengan siempre la forma que esperan los componentes,
// aunque el backend sea una versión anterior (p. ej. features como texto).
function normalizeContent(data) {
  const list = (v) => (Array.isArray(v) ? v : []);
  return {
    ...data,
    services: list(data.services).map((s) => ({
      ...s,
      features: list(s.features).map((f) =>
        typeof f === 'string' ? { title: f, description: '', includes: [] } : { ...f, includes: list(f.includes) },
      ),
      technologies: list(s.technologies),
      benefits: list(s.benefits),
      idealFor: list(s.idealFor),
      faq: list(s.faq),
    })),
  };
}

export async function sendContact(data) {
  let res;
  try {
    res = await fetch(`${API_BASE}/api/contact`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
  } catch {
    throw new Error('No hay conexión con el servidor. Intenta de nuevo.');
  }

  const body = await res.json().catch(() => ({}));
  if (!res.ok) {
    const err = new Error(body.error || 'No se pudo enviar el mensaje.');
    err.fields = body.fields || {};
    throw err;
  }
  return body;
}
