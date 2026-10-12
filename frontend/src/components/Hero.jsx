import { useState } from 'react';
import Icon from './Icon.jsx';

const prefersReducedMotion =
  typeof window !== 'undefined' && window.matchMedia?.('(prefers-reduced-motion: reduce)').matches;

// Portada en carrusel: una imagen por servicio (public/img/<id>-banner.jpg).
// Avanza sola cuando termina la barra de progreso de la pestaña activa y se
// pausa al pasar el cursor encima.
export default function Hero({ company, services, onRequest }) {
  const [index, setIndex] = useState(0);
  const [paused, setPaused] = useState(false);
  const count = services.length;
  const s = services[index];

  const go = (i) => setIndex((i + count) % count);

  return (
    <section
      id="inicio"
      className={`cover ${paused ? 'is-paused' : ''}`}
      aria-roledescription="carrusel"
      aria-label="Servicios de NGN NetVision Technologies"
      onMouseEnter={() => setPaused(true)}
      onMouseLeave={() => setPaused(false)}
    >
      <div className="cover__slides">
        {services.map((item, i) => (
          <div
            key={item.id}
            className={`cover__slide cover__slide--${item.id} ${i === index ? 'is-active' : ''}`}
            aria-hidden={i !== index}
          >
            <img src={`/img/${item.id}-banner.jpg`} alt="" className="cover__img" />
          </div>
        ))}
      </div>

      {/* key: al cambiar de servicio los textos vuelven a entrar con animación */}
      <div key={s.id} className={`container cover__content accent--${s.id}`}>
        <span className="cover__brand">
          <span className="brand__ngn">NGN</span>
          <span className="brand__net">NetVision</span>
          <span className="cover__brand-tech">Technologies</span>
        </span>
        <h1 className="cover__title">
          <span className="cover__title-top">{s.title}</span>
          <span className="cover__title-main">{s.subtitle}</span>
        </h1>
        <p className="cover__lead">{s.summary}</p>
        <div className="cover__actions">
          <button type="button" className="btn btn--primary btn--lg" onClick={() => onRequest(s.id)}>
            Solicitar cotización <Icon name="arrow-right" size={18} />
          </button>
          <a href="#servicios" className="btn btn--glass btn--lg">
            Ver todos los servicios
          </a>
        </div>
      </div>

      <button type="button" className="cover__arrow cover__arrow--prev" onClick={() => go(index - 1)} aria-label="Servicio anterior">
        <Icon name="arrow-right" size={22} />
      </button>
      <button type="button" className="cover__arrow cover__arrow--next" onClick={() => go(index + 1)} aria-label="Servicio siguiente">
        <Icon name="arrow-right" size={22} />
      </button>

      <div className="cover__tabs">
        {services.map((item, i) => (
          <button
            key={item.id}
            type="button"
            className={`cover__tab accent--${item.id} ${i === index ? 'is-active' : ''}`}
            onClick={() => go(i)}
            aria-current={i === index}
            aria-label={`Ver ${item.title}`}
          >
            <span className="icon-badge">
              <Icon name={item.icon} size={20} />
            </span>
            <span className="cover__tab-text">
              <strong>{item.title}</strong>
              <small>{item.subtitle}</small>
            </span>
            {i === index && (
              <span
                className={`cover__progress ${prefersReducedMotion ? 'is-static' : ''}`}
                onAnimationEnd={() => !prefersReducedMotion && go(index + 1)}
              />
            )}
          </button>
        ))}
      </div>
      <span className="sr-only" aria-live="polite">
        {company.name}: {s.title}
      </span>
    </section>
  );
}
