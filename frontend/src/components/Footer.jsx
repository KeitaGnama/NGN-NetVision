import Icon, { Logo } from './Icon.jsx';

export default function Footer({ company, services }) {
  return (
    <footer className="footer">
      <div className="container footer__showcase">
        {services.map((s) => (
          <a key={s.id} href="#servicios" className={`footer__tile accent--${s.id}`}>
            <img src={`/img/${s.id}-banner.jpg`} alt="" loading="lazy" />
            <span className="footer__tile-label">
              <span className="icon-badge">
                <Icon name={s.icon} size={18} />
              </span>
              {s.title}
            </span>
          </a>
        ))}
      </div>

      <div className="container footer__grid">
        <div className="footer__brand">
          <a href="#inicio" className="brand brand--light">
            <Logo size={42} />
            <span className="brand__name">
              <span className="brand__line">
                <span className="brand__ngn">NGN</span>
                <span className="brand__net">NetVision</span>
              </span>
              <span className="brand__tech">Technologies</span>
            </span>
          </a>
          <p>{company.description}</p>
        </div>

        <div>
          <h5>Servicios</h5>
          <ul>
            {services.map((s) => (
              <li key={s.id}>
                <a href={`#${s.id}`}>{s.title}</a>
              </li>
            ))}
          </ul>
        </div>

        <div>
          <h5>Empresa</h5>
          <ul>
            <li><a href="#nosotros">Nosotros</a></li>
            <li><a href="#proceso">Cómo trabajamos</a></li>
            <li><a href="#contacto">Contacto</a></li>
          </ul>
        </div>

        <div>
          <h5>Contacto</h5>
          <ul>
            <li><a href={`mailto:${company.email}`}>{company.email}</a></li>
            <li><a href={`tel:${company.phone.replace(/\s/g, '')}`}>{company.phone}</a></li>
            <li>{company.address}</li>
          </ul>
        </div>
      </div>
      <div className="container footer__bottom">
        <span>© {new Date().getFullYear()} {company.name}. Todos los derechos reservados.</span>
        <span>{services.map((s) => s.title).join(' · ')}</span>
      </div>

      <div className="footer__wordmark" aria-hidden="true">
        <span className="brand__ngn">NGN</span> <span className="brand__net">NetVision</span>
      </div>
    </footer>
  );
}
