import { useEffect, useRef, useState } from 'react';
import Icon from './Icon.jsx';

// Ventana que se abre al dar clic en "Ver detalles" de un servicio.
export default function ServicePanel({ service: s, onClose, onRequest }) {
  const closeRef = useRef(null);
  const [feature, setFeature] = useState(null);
  const featureRef = useRef(null);
  featureRef.current = feature;

  useEffect(() => {
    const previousFocus = document.activeElement;
    closeRef.current?.focus();
    document.body.style.overflow = 'hidden';

    // Esc cierra primero la ventana del elemento, y luego la del servicio.
    const onKey = (e) => {
      if (e.key !== 'Escape') return;
      if (featureRef.current) setFeature(null);
      else onClose();
    };
    window.addEventListener('keydown', onKey);

    return () => {
      window.removeEventListener('keydown', onKey);
      document.body.style.overflow = '';
      previousFocus?.focus?.({ preventScroll: true });
    };
  }, [onClose]);

  const request = () => {
    onClose();
    onRequest(s.id);
  };

  return (
    <div className={`service-panel ${feature ? 'has-feature' : ''}`} onClick={onClose}>
      <aside
        className={`service-panel__drawer accent--${s.id}`}
        role="dialog"
        aria-modal="true"
        aria-labelledby="service-panel-title"
        onClick={(e) => e.stopPropagation()}
      >
        <header className="service-panel__header">
          <span className="icon-badge icon-badge--lg">
            <Icon name={s.icon} size={28} />
          </span>
          <div className="service-panel__heading">
            <span className="service-panel__subtitle">{s.subtitle}</span>
            <h3 id="service-panel-title">{s.title}</h3>
          </div>
          <button ref={closeRef} className="service-panel__close" onClick={onClose} aria-label="Cerrar detalles">
            <Icon name="x" size={20} />
          </button>
        </header>

        <div className="service-panel__body">
          <p className="service-panel__lead">{s.description}</p>

          <PanelSection title="¿Qué incluye?">
            <ul className="check-list">
              {s.features.map((f) => (
                <li key={f.title}>
                  <button
                    type="button"
                    className={`feature-btn ${feature?.title === f.title ? 'feature-btn--active' : ''}`}
                    onClick={() => setFeature(f)}
                    aria-haspopup="dialog"
                    aria-expanded={feature?.title === f.title}
                  >
                    <span className="check-list__mark">
                      <Icon name="check" size={14} strokeWidth={3} />
                    </span>
                    <span className="feature-btn__text">{f.title}</span>
                    <Icon name="arrow-right" size={16} className="feature-btn__arrow" />
                  </button>
                </li>
              ))}
            </ul>
            <p className="service-panel__hint">Da clic en cada elemento para ver más detalles.</p>
          </PanelSection>

          {s.benefits?.length > 0 && (
            <PanelSection title="Beneficios para tu empresa">
              <ul className="benefit-list">
                {s.benefits.map((b) => (
                  <li key={b}>
                    <Icon name="zap" size={16} />
                    {b}
                  </li>
                ))}
              </ul>
            </PanelSection>
          )}

          {s.idealFor?.length > 0 && (
            <PanelSection title="Ideal para">
              <div className="tags">
                {s.idealFor.map((t) => (
                  <span key={t} className="tag tag--accent">
                    {t}
                  </span>
                ))}
              </div>
            </PanelSection>
          )}

          <PanelSection title="Tecnologías">
            <div className="tags">
              {s.technologies.map((t) => (
                <span key={t} className="tag">
                  {t}
                </span>
              ))}
            </div>
          </PanelSection>

          {s.faq?.length > 0 && (
            <PanelSection title="Preguntas frecuentes">
              <div className="faq">
                {s.faq.map((q) => (
                  <details key={q.question} className="faq__item">
                    <summary>{q.question}</summary>
                    <p>{q.answer}</p>
                  </details>
                ))}
              </div>
            </PanelSection>
          )}
        </div>

        <footer className="service-panel__footer">
          <button className="btn btn--accent btn--block" onClick={request}>
            Solicitar información <Icon name="arrow-right" size={18} />
          </button>
        </footer>
      </aside>

      {feature && (
        <FeatureModal
          key={feature.title}
          service={s}
          feature={feature}
          onBack={() => setFeature(null)}
          onRequest={(message) => {
            onClose();
            onRequest(s.id, message);
          }}
        />
      )}
    </div>
  );
}

// Segunda ventana: detalle de un elemento de "¿Qué incluye?". Sus opciones se
// pueden seleccionar y se envían ya escritas al formulario de contacto.
function FeatureModal({ service: s, feature: f, onBack, onRequest }) {
  const backRef = useRef(null);
  const [selected, setSelected] = useState([]);
  const [otherOn, setOtherOn] = useState(false);
  const [otherText, setOtherText] = useState('');
  const otherRef = useRef(null);

  const toggle = (option) =>
    setSelected((list) => (list.includes(option) ? list.filter((o) => o !== option) : [...list, option]));

  const toggleOther = () => {
    setOtherOn((on) => !on);
    if (!otherOn) setTimeout(() => otherRef.current?.focus(), 0);
  };

  const total = selected.length + (otherOn ? 1 : 0);

  const quote = () => {
    const chosen = f.includes.filter((o) => selected.includes(o));
    if (otherOn) chosen.push(`Otro: ${otherText.trim() || '(por especificar)'}`);
    const lines = [`Me interesa: ${s.title} — ${f.title}.`];
    if (chosen.length > 0) lines.push('', 'Opciones que necesito:', ...chosen.map((o) => `• ${o}`));
    lines.push('', 'Detalles de mi proyecto: ');
    onRequest(lines.join('\n'));
  };

  useEffect(() => {
    const previousFocus = document.activeElement;
    backRef.current?.focus();
    return () => previousFocus?.focus?.({ preventScroll: true });
  }, []);

  return (
    // Se muestra junto a la ventana del servicio (a la derecha), no encima.
    <aside
      className={`service-panel__drawer feature-modal accent--${s.id}`}
      role="dialog"
      aria-modal="true"
      aria-labelledby="feature-modal-title"
      onClick={(e) => e.stopPropagation()}
    >
        <header className="feature-modal__header">
          <span className="feature-modal__crumb">{s.title}</span>
          <button ref={backRef} type="button" className="service-panel__close feature-modal__close" onClick={onBack} aria-label="Cerrar detalle">
            <Icon name="x" size={18} />
          </button>
          <span className="icon-badge">
            <Icon name="check" size={22} strokeWidth={2.5} />
          </span>
          <h3 id="feature-modal-title">{f.title}</h3>
          {f.time && (
            <span className="feature-modal__time">
              <Icon name="clock" size={15} /> Tiempo estimado: <strong>{f.time}</strong>
            </span>
          )}
        </header>

        <div className="service-panel__body">
          {f.description && <p className="service-panel__lead">{f.description}</p>}

          {f.includes.length > 0 && (
            <PanelSection title="Elige lo que necesitas">
              <div className="option-list">
                {f.includes.map((option) => {
                  const on = selected.includes(option);
                  return (
                    <button
                      key={option}
                      type="button"
                      className={`option ${on ? 'option--on' : ''}`}
                      aria-pressed={on}
                      onClick={() => toggle(option)}
                    >
                      <span className="option__box">
                        <Icon name="check" size={14} strokeWidth={3} />
                      </span>
                      {option}
                    </button>
                  );
                })}

                <button
                  type="button"
                  className={`option ${otherOn ? 'option--on' : ''}`}
                  aria-pressed={otherOn}
                  aria-expanded={otherOn}
                  onClick={toggleOther}
                >
                  <span className="option__box">
                    <Icon name="check" size={14} strokeWidth={3} />
                  </span>
                  Otro
                </button>
                {otherOn && (
                  <input
                    ref={otherRef}
                    className="option__other"
                    value={otherText}
                    onChange={(e) => setOtherText(e.target.value)}
                    placeholder="Escribe qué otra cosa necesitas"
                    maxLength={200}
                    aria-label="Describe la otra opción que necesitas"
                  />
                )}
              </div>
              <p className="service-panel__hint">
                {total === 0
                  ? 'Selecciona una o varias opciones para tu cotización.'
                  : `${total} ${total === 1 ? 'opción seleccionada' : 'opciones seleccionadas'}`}
              </p>
            </PanelSection>
          )}
        </div>

        <footer className="service-panel__footer feature-modal__footer">
          <button type="button" className="btn btn--outline" onClick={onBack}>
            Cerrar
          </button>
          <button type="button" className="btn btn--accent" onClick={quote}>
            Cotizar{total > 0 ? ` (${total})` : ''} <Icon name="arrow-right" size={18} />
          </button>
        </footer>
    </aside>
  );
}

function PanelSection({ title, children }) {
  return (
    <section className="service-panel__section">
      <h4>{title}</h4>
      {children}
    </section>
  );
}
