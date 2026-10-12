import { useCallback, useEffect, useState } from 'react';
import { getContent } from './api.js';
import { useReveal } from './hooks/useReveal.js';
import Icon, { Logo } from './components/Icon.jsx';
import Navbar from './components/Navbar.jsx';
import Hero from './components/Hero.jsx';
import Services from './components/Services.jsx';
import About from './components/About.jsx';
import Process from './components/Process.jsx';
import Providers from './components/Providers.jsx';
import Contact from './components/Contact.jsx';
import Footer from './components/Footer.jsx';
import WhatsAppButton from './components/WhatsAppButton.jsx';

export default function App() {
  const [content, setContent] = useState(null);
  const [status, setStatus] = useState('loading');
  const [selectedService, setSelectedService] = useState('');
  const [draft, setDraft] = useState(null); // mensaje prellenado desde "Cotizar"

  const load = useCallback(() => {
    setStatus('loading');
    getContent()
      .then((data) => {
        setContent(data);
        setStatus('ready');
      })
      .catch(() => setStatus('error'));
  }, []);

  useEffect(() => {
    load();
  }, [load]);

  useReveal(status === 'ready');

  if (status === 'loading') {
    return (
      <div className="page-state">
        <Logo size={56} />
        <div className="spinner" aria-label="Cargando" />
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="page-state">
        <Icon name="alert" size={48} className="page-state__icon" />
        <h1>No pudimos cargar el sitio</h1>
        <p>Verifica que el servidor esté en ejecución e intenta de nuevo.</p>
        <button className="btn btn--primary" onClick={load}>
          Reintentar
        </button>
      </div>
    );
  }

  const { company, services, highlights, process } = content;

  const requestService = (id, message) => {
    setSelectedService(id);
    if (message) setDraft({ text: message });
    document.getElementById('contacto')?.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <>
      <Navbar />
      <main>
        <Hero company={company} services={services} onRequest={requestService} />
        <Services services={services} onRequest={requestService} />
        <About company={company} highlights={highlights} />
        <Process steps={process} />
        <Providers services={services} />
        <Contact
          company={company}
          services={services}
          selectedService={selectedService}
          draft={draft}
          onServiceChange={setSelectedService}
        />
      </main>
      <Footer company={company} services={services} />
      <WhatsAppButton number={company.whatsapp} company={company.name} />
    </>
  );
}
