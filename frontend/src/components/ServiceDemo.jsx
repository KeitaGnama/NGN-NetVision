import { useEffect, useRef, useState } from 'react';
import Icon from './Icon.jsx';

// "Video" animado que simula cada servicio. No es un archivo de video: se
// dibuja en tiempo real (más liviano y nítido) y se controla como uno, con
// reproducir/pausa, barra de progreso y tiempo. Solo avanza cuando está a la vista.

const DURATION = { software: 12000, redes: 10000, tracking: 12000 };
const TITLES = {
  software: 'Desarrollo de una app web',
  redes: 'Monitoreo de red empresarial',
  tracking: 'Rastreo de flotilla en vivo',
};

const prefersReducedMotion =
  typeof window !== 'undefined' && window.matchMedia?.('(prefers-reduced-motion: reduce)').matches;

export default function ServiceDemo({ service }) {
  const id = service.id;
  const duration = DURATION[id] ?? 10000;
  const ref = useRef(null);
  const [playing, setPlaying] = useState(!prefersReducedMotion);
  const [visible, setVisible] = useState(false);
  const [t, setT] = useState(prefersReducedMotion ? duration * 0.8 : 0);

  useEffect(() => {
    const el = ref.current;
    if (!el) return undefined;
    const io = new IntersectionObserver(([e]) => setVisible(e.isIntersecting), { threshold: 0.25 });
    io.observe(el);
    return () => io.disconnect();
  }, []);

  useEffect(() => {
    if (!playing || !visible) return undefined;
    let raf;
    let last = performance.now();
    const tick = (now) => {
      setT((prev) => (prev + (now - last)) % duration);
      last = now;
      raf = requestAnimationFrame(tick);
    };
    raf = requestAnimationFrame(tick);
    return () => cancelAnimationFrame(raf);
  }, [playing, visible, duration]);

  const seek = (e) => {
    const r = e.currentTarget.getBoundingClientRect();
    setT(((e.clientX - r.left) / r.width) * duration);
  };

  return (
    <div ref={ref} className={`demo accent--${id}`}>
      <div className="demo__bar">
        <span className="demo__dots" aria-hidden="true">
          <i />
          <i />
          <i />
        </span>
        <span className="demo__title">{TITLES[id] ?? service.title}</span>
        <span className="demo__badge">Simulación</span>
      </div>

      <div className="demo__stage" role="img" aria-label={`Simulación animada: ${TITLES[id] ?? service.title}`}>
        {id === 'software' && <SoftwareScene t={t} />}
        {id === 'redes' && <NetworkScene t={t} />}
        {id === 'tracking' && <TrackingScene t={t} duration={duration} />}
      </div>

      <div className="demo__controls">
        <button
          type="button"
          className="demo__play"
          onClick={() => setPlaying((p) => !p)}
          aria-label={playing ? 'Pausar simulación' : 'Reproducir simulación'}
        >
          <Icon name={playing ? 'pause' : 'play'} size={16} />
        </button>
        <div className="demo__track" onClick={seek} aria-hidden="true">
          <span className="demo__progress" style={{ width: `${(t / duration) * 100}%` }} />
        </div>
        <span className="demo__time">
          {fmt(t)} / {fmt(duration)}
        </span>
      </div>
    </div>
  );
}

const fmt = (ms) => `0:${String(Math.floor(ms / 1000)).padStart(2, '0')}`;
const clamp = (v, a, b) => Math.min(b, Math.max(a, v));

/* ---------- Software: se escribe el código, se compila y aparece la app ---------- */

const CODE = [
  [['k', 'const '], ['v', 'app'], ['p', ' = '], ['f', 'crearApp'], ['p', '({']],
  [['p', '  '], ['v', 'nombre'], ['p', ': '], ['s', "'Mi Negocio'"], ['p', ',']],
  [['p', '  '], ['v', 'modulos'], ['p', ': ['], ['s', "'ventas'"], ['p', ', '], ['s', "'stock'"], ['p', '],']],
  [['p', '});']],
  [],
  [['v', 'app'], ['p', '.'], ['f', 'conectar'], ['p', '('], ['s', "'/api/pagos'"], ['p', ');']],
  [['k', 'await '], ['v', 'app'], ['p', '.'], ['f', 'publicar'], ['p', '();']],
];
const LINE_LEN = CODE.map((l) => l.reduce((n, [, x]) => n + x.length, 0));
const LINE_START = LINE_LEN.map((_, i) => LINE_LEN.slice(0, i).reduce((n, len) => n + len + 1, 0));
const TOTAL_CHARS = LINE_START.at(-1) + LINE_LEN.at(-1);

function SoftwareScene({ t }) {
  const typed = Math.max(0, (t - 400) * 0.034);
  const done = typed >= TOTAL_CHARS;
  const caretLine = done ? CODE.length - 1 : LINE_START.findLastIndex((s) => s <= typed);

  return (
    <div className="sw">
      <div className="sw__editor">
        <div className="sw__tabs">
          <span className="is-active">app.js</span>
          <span>estilos.css</span>
        </div>
        <pre className="sw__code">
          {CODE.map((tokens, i) => {
            let left = clamp(typed - LINE_START[i], 0, LINE_LEN[i]);
            return (
              <div key={i} className="sw__line">
                <span className="sw__ln">{i + 1}</span>
                {tokens.map(([cls, text], j) => {
                  const part = text.slice(0, left);
                  left = Math.max(0, left - text.length);
                  return part ? (
                    <span key={j} className={`tok-${cls}`}>
                      {part}
                    </span>
                  ) : null;
                })}
                {i === caretLine && <span className="sw__caret" />}
              </div>
            );
          })}
        </pre>
        <div className="sw__terminal">
          {t > 6000 && <div className="sw__term-line">$ npm run build</div>}
          {t > 7000 && <div className="sw__term-line is-ok">✓ Compilado en 2.1 s</div>}
          {t > 7900 && <div className="sw__term-line is-ok">✓ Publicado en tunegocio.com</div>}
        </div>
      </div>

      <div className="sw__preview">
        <div className="sw__browser">
          <span />
          <span />
          <span />
          <em>tunegocio.com</em>
        </div>
        <div className={`sw__app ${t > 8200 ? 'is-live' : ''}`}>
          {t > 8200 ? (
            <>
              <div className="sw__hero pop" />
              <div className="sw__cards">
                <div className="pop" style={{ animationDelay: '0.15s' }} />
                <div className="pop" style={{ animationDelay: '0.3s' }} />
                <div className="pop" style={{ animationDelay: '0.45s' }} />
              </div>
              <div className="sw__chart pop" style={{ animationDelay: '0.6s' }}>
                {[40, 65, 50, 80, 70, 95].map((h, i) => (
                  <i key={i} style={{ height: `${h}%` }} />
                ))}
              </div>
            </>
          ) : (
            <div className="sw__waiting">
              <span className="sw__spinner" />
              Esperando despliegue…
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

/* ---------- Redes: paquetes viajando, firewall bloquea una amenaza ---------- */

const NODES = {
  cloud: { x: 200, y: 36, icon: 'cloud', label: 'Internet' },
  fw: { x: 200, y: 100, icon: 'shield', label: 'Firewall' },
  sw: { x: 200, y: 162, icon: 'network', label: 'Switch' },
  pc: { x: 62, y: 222, icon: 'monitor', label: 'Oficina' },
  cam: { x: 152, y: 226, icon: 'camera', label: 'CCTV' },
  tel: { x: 248, y: 226, icon: 'phone', label: 'Telefonía IP' },
  srv: { x: 338, y: 222, icon: 'server', label: 'Servidor' },
};
const LINKS = [['cloud', 'fw'], ['fw', 'sw'], ['sw', 'pc'], ['sw', 'cam'], ['sw', 'tel'], ['sw', 'srv']];

function NetworkScene({ t }) {
  const threat = t > 4800 && t < 7800;
  const blocked = t > 5600 && t < 7800;
  const latency = 3 + Math.round(Math.abs(Math.sin(t / 900)) * 3);

  return (
    <div className="net">
      <svg viewBox="0 0 400 260" className="net__svg" aria-hidden="true">
        {LINKS.map(([a, b]) => (
          <line key={a + b} x1={NODES[a].x} y1={NODES[a].y} x2={NODES[b].x} y2={NODES[b].y} className="net__link" />
        ))}
        {LINKS.flatMap(([a, b], k) =>
          [0, 0.5].map((phase) => {
            let f = ((t / 1600 + phase + k * 0.13) % 1);
            if (k % 2) f = 1 - f;
            const x = NODES[a].x + (NODES[b].x - NODES[a].x) * f;
            const y = NODES[a].y + (NODES[b].y - NODES[a].y) * f;
            return <circle key={`${a}${b}${phase}`} cx={x} cy={y} r="3.6" className="net__packet" />;
          }),
        )}
        {threat && (
          <circle
            cx={NODES.cloud.x}
            cy={NODES.cloud.y + (NODES.fw.y - NODES.cloud.y - 14) * clamp((t - 4800) / 800, 0, 1)}
            r="5"
            className="net__threat"
          />
        )}
      </svg>

      {Object.entries(NODES).map(([key, n]) => (
        <div
          key={key}
          className={`net__node ${key === 'fw' && blocked ? 'is-alert' : ''}`}
          style={{ left: `${n.x / 4}%`, top: `${n.y / 2.6}%` }}
        >
          <span className="net__icon">
            <Icon name={n.icon} size={16} />
          </span>
          <small>{n.label}</small>
        </div>
      ))}

      <div className="net__hud">
        <span>
          Latencia <b>{latency} ms</b>
        </span>
        <span>
          Disponibilidad <b>99.98%</b>
        </span>
      </div>

      <div className="net__bars" aria-hidden="true">
        {Array.from({ length: 12 }, (_, i) => (
          <i key={i} style={{ height: `${25 + 70 * Math.abs(Math.sin(t / 650 + i * 0.55))}%` }} />
        ))}
      </div>

      {blocked && (
        <div className="demo__toast is-danger">
          <Icon name="shield" size={14} /> Amenaza bloqueada por el firewall
        </div>
      )}
      {!threat && t > 7800 && t < 9800 && (
        <div className="demo__toast is-ok">
          <Icon name="check" size={14} strokeWidth={3} /> Red segura y estable
        </div>
      )}
    </div>
  );
}

/* ---------- GPS: la camioneta recorre la ruta, geocerca y llegada ---------- */

const ROUTE = [[34, 228], [34, 150], [128, 150], [128, 82], [250, 82], [250, 150], [362, 150], [362, 46]];
const SEG = ROUTE.slice(1).map((p, i) => Math.hypot(p[0] - ROUTE[i][0], p[1] - ROUTE[i][1]));
const ROUTE_LEN = SEG.reduce((a, b) => a + b, 0);
const FENCE = { x: 250, y: 116, r: 40 };

function pointAt(d) {
  let rest = d;
  for (let i = 0; i < SEG.length; i += 1) {
    if (rest <= SEG[i]) {
      const f = SEG[i] ? rest / SEG[i] : 0;
      const [x1, y1] = ROUTE[i];
      const [x2, y2] = ROUTE[i + 1];
      return { x: x1 + (x2 - x1) * f, y: y1 + (y2 - y1) * f, i };
    }
    rest -= SEG[i];
  }
  const [x, y] = ROUTE.at(-1);
  return { x, y, i: SEG.length - 1 };
}

function TrackingScene({ t, duration }) {
  const p = clamp(t / (duration * 0.82), 0, 1);
  const pos = pointAt(p * ROUTE_LEN);
  const arrived = p >= 1;
  const inFence = Math.hypot(pos.x - FENCE.x, pos.y - FENCE.y) < FENCE.r;
  const speed = arrived ? 0 : Math.round(54 + 16 * Math.sin(t / 650));
  const traveled = [...ROUTE.slice(0, pos.i + 1), [pos.x, pos.y]].map((q) => q.join(',')).join(' ');

  return (
    <div className="gps">
      <svg viewBox="0 0 400 260" className="gps__svg" aria-hidden="true">
        <rect width="400" height="260" className="gps__bg" />
        <path d="M0 200 C 90 180, 140 240, 220 210 S 360 190, 400 220" className="gps__river" />
        <rect x="290" y="190" width="90" height="50" rx="8" className="gps__park" />
        <rect x="60" y="20" width="50" height="40" rx="6" className="gps__park" />
        {[34, 128, 250, 362].map((x) => (
          <line key={`v${x}`} x1={x} y1="0" x2={x} y2="260" className="gps__street" />
        ))}
        {[46, 82, 150, 228].map((y) => (
          <line key={`h${y}`} x1="0" y1={y} x2="400" y2={y} className="gps__street" />
        ))}
        <circle cx={FENCE.x} cy={FENCE.y} r={FENCE.r} className={`gps__fence ${inFence ? 'is-in' : ''}`} />
        <polyline points={ROUTE.map((q) => q.join(',')).join(' ')} className="gps__route" />
        <polyline points={traveled} className="gps__traveled" />
      </svg>

      <div className="gps__pin" style={{ left: `${ROUTE.at(-1)[0] / 4}%`, top: `${ROUTE.at(-1)[1] / 2.6}%` }}>
        <Icon name="map-pin" size={20} />
      </div>
      <div className="gps__van" style={{ left: `${pos.x / 4}%`, top: `${pos.y / 2.6}%` }}>
        <Icon name="truck" size={14} />
      </div>

      <div className="gps__card">
        <span className="gps__card-head">
          <strong>Vehículo #12</strong>
          <span className="gps__signal" aria-label="Señal GPS">
            <i />
            <i />
            <i />
          </span>
        </span>
        <span className={`gps__status ${arrived ? 'is-stop' : ''}`}>{arrived ? 'Detenido' : 'En movimiento'}</span>
        <span className="gps__speed">
          <b>{speed}</b> km/h
        </span>
      </div>

      {inFence && (
        <div className="demo__toast is-warn">
          <Icon name="alert" size={14} /> Entró a geocerca: Almacén Norte
        </div>
      )}
      {arrived && (
        <div className="demo__toast is-ok">
          <Icon name="check" size={14} strokeWidth={3} /> Llegó a destino a tiempo
        </div>
      )}
    </div>
  );
}
