import { useCallback, useState } from 'react';
import Icon from './Icon.jsx';
import SectionHeader from './SectionHeader.jsx';
import ServicePanel from './ServicePanel.jsx';
import ServiceDemo from './ServiceDemo.jsx';

export default function Services({ services, onRequest }) {
  const [openId, setOpenId] = useState(null);
  const closePanel = useCallback(() => setOpenId(null), []);
  const openService = services.find((s) => s.id === openId);

  return (
    <section id="servicios" className="section">
      <div className="container">
        <SectionHeader
          eyebrow="Nuestros servicios"
          title="Tres áreas de especialidad, una sola visión"
          text="Cubrimos lo que tu empresa necesita para operar en la era digital: el software que la impulsa, la red que la conecta y la tecnología que protege tus activos en movimiento."
        />

        <div className="services-grid">
          {services.map((s, i) => (
            <button
              key={s.id}
              type="button"
              className={`service-card accent--${s.id} reveal`}
              style={{ transitionDelay: `${i * 100}ms` }}
              onClick={() => setOpenId(s.id)}
              aria-haspopup="dialog"
            >
              <span className="icon-badge icon-badge--lg">
                <Icon name={s.icon} size={28} />
              </span>
              <h3>{s.title}</h3>
              <p>{s.summary}</p>
              <span className="service-card__more">
                Ver detalles <Icon name="arrow-right" size={16} />
              </span>
            </button>
          ))}
        </div>

        <div className="service-details">
          {services.map((s, i) => (
            <ServiceDetail key={s.id} service={s} reversed={i % 2 === 1} onRequest={onRequest} />
          ))}
        </div>
      </div>

      {openService && <ServicePanel service={openService} onClose={closePanel} onRequest={onRequest} />}
    </section>
  );
}

function ServiceDetail({ service: s, reversed, onRequest }) {
  return (
    <article id={s.id} className={`service-detail accent--${s.id} ${reversed ? 'service-detail--reversed' : ''}`}>
      <div className="service-detail__content reveal">
        <span className="service-detail__label">
          <Icon name={s.icon} size={18} /> {s.subtitle}
        </span>
        <h3>{s.title}</h3>
        <p>{s.description}</p>
        <ul className="check-list">
          {s.features.map((f) => (
            <li key={f.title}>
              <span className="check-list__mark">
                <Icon name="check" size={14} strokeWidth={3} />
              </span>
              {f.title}
            </li>
          ))}
        </ul>
        <div className="tags">
          {s.technologies.map((t) => (
            <span key={t} className="tag">
              {t}
            </span>
          ))}
        </div>
        <button className="btn btn--accent" onClick={() => onRequest(s.id)}>
          Solicitar información <Icon name="arrow-right" size={18} />
        </button>
      </div>

      <div className="service-detail__visual reveal">
        <ServiceDemo service={s} />
      </div>
    </article>
  );
}
