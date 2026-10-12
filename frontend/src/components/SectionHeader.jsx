export default function SectionHeader({ eyebrow, title, text, align = 'center' }) {
  return (
    <div className={`section-header section-header--${align} reveal`}>
      <span className="eyebrow">{eyebrow}</span>
      <h2>{title}</h2>
      {text && <p>{text}</p>}
    </div>
  );
}
