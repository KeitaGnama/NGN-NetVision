import { useEffect, useState } from 'react';
import Icon, { Logo } from './Icon.jsx';

const links = [
  { href: '#inicio', label: 'Inicio' },
  { href: '#servicios', label: 'Servicios' },
  { href: '#nosotros', label: 'Nosotros' },
  { href: '#proceso', label: 'Proceso' },
  { href: '#contacto', label: 'Contacto' },
];

export default function Navbar() {
  const [open, setOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 10);
    onScroll();
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  useEffect(() => {
    document.body.style.overflow = open ? 'hidden' : '';
    const onKey = (e) => e.key === 'Escape' && setOpen(false);
    if (open) window.addEventListener('keydown', onKey);
    return () => {
      document.body.style.overflow = '';
      window.removeEventListener('keydown', onKey);
    };
  }, [open]);

  const close = () => setOpen(false);

  return (
    <header className={`navbar ${scrolled || open ? 'navbar--scrolled' : ''}`}>
      <div className="container navbar__inner">
        <a href="#inicio" className="brand" onClick={close}>
          <Logo size={42} />
          <span className="brand__name">
            <span className="brand__line">
              <span className="brand__ngn">NGN</span>
              <span className="brand__net">NetVision</span>
            </span>
            <span className="brand__tech">Technologies</span>
          </span>
        </a>

        {open && <div className="navbar__backdrop" onClick={close} aria-hidden="true" />}

        <nav className={`navbar__nav ${open ? 'is-open' : ''}`} aria-label="Principal">
          {links.map((link) => (
            <a key={link.href} href={link.href} className="navbar__link" onClick={close}>
              {link.label}
            </a>
          ))}
          <a href="#contacto" className="btn btn--primary btn--sm navbar__cta" onClick={close}>
            Solicitar cotización
          </a>
        </nav>

        <button
          className="navbar__toggle"
          aria-label={open ? 'Cerrar menú' : 'Abrir menú'}
          aria-expanded={open}
          onClick={() => setOpen((o) => !o)}
        >
          <Icon name={open ? 'x' : 'menu'} />
        </button>
      </div>
    </header>
  );
}
