import Icon from './Icon.jsx';
import SectionHeader from './SectionHeader.jsx';

// "Nuestros proveedores": una columna por servicio. Los proveedores se editan
// en backend/content.go (campo Providers de cada servicio).
export default function Providers({ services }) {
  return (
    <section id="proveedores" className="section providers-section">
      <div className="container">
        <SectionHeader
          eyebrow="Nuestros proveedores"
          title="Trabajamos con marcas líderes"
          text="Usamos equipos y plataformas de fabricantes reconocidos para que cada solución sea confiable, segura y con respaldo."
        />

        <div className="providers">
          {services.map((s, i) => (
            <div
              key={s.id}
              className={`providers__group accent--${s.id} reveal`}
              style={{ transitionDelay: `${i * 120}ms` }}
            >
              <header className="providers__head">
                <span className="icon-badge">
                  <Icon name={s.icon} size={22} />
                </span>
                <div>
                  <h3>{s.title}</h3>
                  <small>
                    {s.providers.length > 0
                      ? `${s.providers.length} ${s.providers.length === 1 ? 'proveedor' : 'proveedores'}`
                      : 'Próximamente'}
                  </small>
                </div>
              </header>

              {s.providers.length > 0 ? (
                <ul className="providers__list">
                  {s.providers.map((p, j) => (
                    <li key={p.name} className="provider" style={{ animationDelay: `${j * 60}ms` }}>
                      <span className="provider__mark" aria-hidden="true">
                        {initials(p.name)}
                      </span>
                      <span className="provider__text">
                        <strong>{p.name}</strong>
                        {p.description && <small>{p.description}</small>}
                      </span>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="providers__soon">
                  <Icon name="clock" size={22} />
                  <p>Muy pronto daremos a conocer a nuestros proveedores de este servicio.</p>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

// Monograma con las iniciales de la marca (p. ej. "MikroTik" → "MT", "AWS" → "AW").
function initials(name) {
  const caps = name.match(/[A-Z0-9]/g) ?? [];
  if (caps.length >= 2) return caps.slice(0, 2).join('');
  return name.replace(/[^A-Za-z0-9]/g, '').slice(0, 2).toUpperCase();
}
