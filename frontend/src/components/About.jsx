import Icon from './Icon.jsx';
import SectionHeader from './SectionHeader.jsx';

export default function About({ company, highlights }) {
  return (
    <section id="nosotros" className="section section--muted">
      <div className="container about">
        <div className="about__intro">
          <SectionHeader eyebrow="Nosotros" title="Tu socio tecnológico de confianza" align="left" />
          <p className="about__text reveal">{company.about}</p>
          <div className="about__mv">
            <div className="mv-card reveal">
              <h4>Misión</h4>
              <p>{company.mission}</p>
            </div>
            <div className="mv-card reveal">
              <h4>Visión</h4>
              <p>{company.vision}</p>
            </div>
          </div>
        </div>

        <div className="highlights">
          {highlights.map((h, i) => (
            <div key={h.title} className="highlight reveal" style={{ transitionDelay: `${i * 80}ms` }}>
              <span className="icon-badge">
                <Icon name={h.icon} size={22} />
              </span>
              <h4>{h.title}</h4>
              <p>{h.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
