import Icon from './Icon.jsx';
import SectionHeader from './SectionHeader.jsx';

export default function Process({ steps }) {
  return (
    <section id="proceso" className="section section--blue">
      <div className="container">
        <SectionHeader
          eyebrow="Cómo trabajamos"
          title="Un proceso claro de principio a fin"
          text="Sabrás en todo momento en qué etapa está tu proyecto y qué sigue."
        />
        <ol className="process">
          {steps.map((step, i) => (
            <li key={step.title} className="process__step reveal" style={{ transitionDelay: `${i * 100}ms` }}>
              <span className="process__icon">
                <Icon name={step.icon} size={24} />
                <span className="process__num">{i + 1}</span>
              </span>
              <h4>{step.title}</h4>
              <p>{step.description}</p>
            </li>
          ))}
        </ol>
      </div>
    </section>
  );
}
