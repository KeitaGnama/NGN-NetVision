import { useEffect, useRef, useState } from 'react';
import { sendContact } from '../api.js';
import Icon from './Icon.jsx';
import SectionHeader from './SectionHeader.jsx';

const emptyForm = { name: '', email: '', phone: '', company: '', service: '', message: '', website: '' };

function validate(form) {
  const errors = {};
  if (!form.name.trim()) errors.name = 'Ingresa tu nombre.';
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email.trim())) errors.email = 'Ingresa un correo electrónico válido.';
  if (form.message.trim().length < 10) errors.message = 'Cuéntanos un poco más (mínimo 10 caracteres).';
  return errors;
}

export default function Contact({ company, services, selectedService, draft, onServiceChange }) {
  const [form, setForm] = useState(emptyForm);
  const [errors, setErrors] = useState({});
  const [status, setStatus] = useState({ type: 'idle', message: '' });
  const messageRef = useRef(null);

  useEffect(() => {
    if (selectedService) setForm((f) => ({ ...f, service: selectedService }));
  }, [selectedService]);

  // Opciones elegidas en la ventana de detalle: se escriben en el mensaje.
  useEffect(() => {
    if (!draft) return;
    setForm((f) => ({ ...f, message: draft.text }));
    setErrors(({ message: _, ...rest }) => rest);
    setStatus({ type: 'idle', message: '' });
    // Deja el cursor al final del mensaje para que la persona siga escribiendo.
    const t = setTimeout(() => {
      const el = messageRef.current;
      if (!el) return;
      el.focus({ preventScroll: true });
      el.setSelectionRange(el.value.length, el.value.length);
    }, 600);
    return () => clearTimeout(t);
  }, [draft]);

  const update = (e) => {
    const { name, value } = e.target;
    setForm((f) => ({ ...f, [name]: value }));
    if (errors[name]) {
      setErrors(({ [name]: _, ...rest }) => rest);
    }
    if (name === 'service') onServiceChange(value);
  };

  const submit = async (e) => {
    e.preventDefault();
    const found = validate(form);
    if (Object.keys(found).length > 0) {
      setErrors(found);
      setStatus({ type: 'idle', message: '' });
      return;
    }

    setStatus({ type: 'sending', message: '' });
    try {
      const res = await sendContact(form);
      setStatus({ type: 'success', message: res.message });
      setForm(emptyForm);
      setErrors({});
      onServiceChange('');
    } catch (err) {
      setErrors(err.fields || {});
      setStatus({ type: 'error', message: err.message });
    }
  };

  const sending = status.type === 'sending';

  return (
    <section id="contacto" className="section section--muted">
      <div className="container contact">
        <div className="contact__info">
          <SectionHeader
            eyebrow="Contacto"
            title="Hablemos de tu próximo proyecto"
            text="Cuéntanos qué necesitas y un especialista te responderá en menos de 24 horas hábiles."
            align="left"
          />
          <ul className="contact__list reveal">
            <ContactItem icon="mail" label="Correo" value={company.email} href={`mailto:${company.email}`} />
            <ContactItem icon="phone" label="Teléfono" value={company.phone} href={`tel:${company.phone.replace(/\s/g, '')}`} />
            <ContactItem icon="map-pin" label="Ubicación" value={company.address} />
            <ContactItem icon="clock" label="Horario" value={company.hours} />
          </ul>
        </div>

        <form className="contact__form reveal" onSubmit={submit} noValidate>
          <div className="form-grid">
            <Field label="Nombre completo *" name="name" error={errors.name}>
              <input id="name" name="name" value={form.name} onChange={update} autoComplete="name" maxLength={100} />
            </Field>
            <Field label="Correo electrónico *" name="email" error={errors.email}>
              <input id="email" name="email" type="email" value={form.email} onChange={update} autoComplete="email" maxLength={150} />
            </Field>
            <Field label="Teléfono" name="phone" error={errors.phone}>
              <input id="phone" name="phone" type="tel" value={form.phone} onChange={update} autoComplete="tel" maxLength={30} />
            </Field>
            <Field label="Empresa" name="company" error={errors.company}>
              <input id="company" name="company" value={form.company} onChange={update} autoComplete="organization" maxLength={120} />
            </Field>
            <Field label="Servicio de interés" name="service" error={errors.service} full>
              <select id="service" name="service" value={form.service} onChange={update}>
                <option value="">Selecciona una opción</option>
                {services.map((s) => (
                  <option key={s.id} value={s.id}>
                    {s.title}
                  </option>
                ))}
                <option value="otro">Otro / varios servicios</option>
              </select>
            </Field>
            <Field label="Mensaje *" name="message" error={errors.message} full>
              <textarea ref={messageRef} id="message" name="message" rows={6} value={form.message} onChange={update} maxLength={3000} placeholder="Describe brevemente tu proyecto o necesidad" />
            </Field>
          </div>

          <input className="hp" name="website" value={form.website} onChange={update} tabIndex={-1} autoComplete="off" aria-hidden="true" />

          {status.type === 'success' && (
            <p className="form-alert form-alert--success" role="status">
              <Icon name="check" size={18} strokeWidth={3} /> {status.message}
            </p>
          )}
          {status.type === 'error' && (
            <p className="form-alert form-alert--error" role="alert">
              <Icon name="alert" size={18} /> {status.message}
            </p>
          )}

          <button type="submit" className="btn btn--primary btn--lg btn--block" disabled={sending}>
            {sending ? 'Enviando…' : 'Enviar mensaje'}
            {!sending && <Icon name="arrow-right" size={18} />}
          </button>
          <p className="form-note">Tus datos se usan únicamente para responder a tu solicitud.</p>
        </form>
      </div>
    </section>
  );
}

function Field({ label, name, error, full, children }) {
  return (
    <div className={`field ${full ? 'field--full' : ''} ${error ? 'field--error' : ''}`}>
      <label htmlFor={name}>{label}</label>
      {children}
      {error && <span className="field__error">{error}</span>}
    </div>
  );
}

function ContactItem({ icon, label, value, href }) {
  return (
    <li className="contact__item">
      <span className="icon-badge">
        <Icon name={icon} size={20} />
      </span>
      <div>
        <span className="contact__label">{label}</span>
        {href ? <a href={href}>{value}</a> : <span>{value}</span>}
      </div>
    </li>
  );
}
