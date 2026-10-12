package main

// Todo el contenido del sitio vive aquí. Edita este archivo para cambiar
// textos, servicios, cifras o datos de contacto sin tocar el frontend.

type Company struct {
	Name string `json:"name"`
	// El título de la portada va en dos líneas: una delgada arriba (TaglineTop)
	// y otra grande en negritas abajo (Tagline).
	TaglineTop  string `json:"taglineTop"`
	Tagline     string `json:"tagline"`
	Description string `json:"description"`
	About       string `json:"about"`
	Mission     string `json:"mission"`
	Vision      string `json:"vision"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	WhatsApp    string `json:"whatsapp"` // Solo dígitos con lada de país, p. ej. 5215512345678
	Address     string `json:"address"`
	Hours       string `json:"hours"`
}

type Stat struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Service struct {
	ID           string    `json:"id"`
	Icon         string    `json:"icon"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Summary      string    `json:"summary"`
	Description  string    `json:"description"`
	Features     []Feature `json:"features"`
	Technologies []string  `json:"technologies"`
	// Información extra que se muestra en el panel de "Ver detalles".
	Benefits []string `json:"benefits"`
	IdealFor []string `json:"idealFor"`
	FAQ      []FAQ    `json:"faq"`
}

// Feature es cada elemento de "¿Qué incluye?". Al darle clic se abre una
// ventana con su descripción, el tiempo estimado y las opciones que el
// cliente puede elegir para su cotización (Includes).
type Feature struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Time        string   `json:"time"`
	Includes    []string `json:"includes"`
}

type FAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Highlight struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Step struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SiteContent struct {
	Company    Company     `json:"company"`
	Stats      []Stat      `json:"stats"`
	Services   []Service   `json:"services"`
	Highlights []Highlight `json:"highlights"`
	Process    []Step      `json:"process"`
}

var siteContent = SiteContent{
	Company: Company{
		Name:        "NGN NetVision Technologies",
		TaglineTop:  "Tecnología que conecta,",
		Tagline:     "protege y hace crecer tu empresa",
		Description: "Desarrollamos software a la medida, diseñamos redes empresariales confiables y monitoreamos tu flotilla en tiempo real. Todo con un solo socio tecnológico.",
		About:       "En NGN NetVision Technologies unimos tres especialidades que toda empresa moderna necesita: software, conectividad y control de activos. Acompañamos a nuestros clientes desde el diagnóstico hasta el soporte continuo, con soluciones escalables y un equipo que habla tu idioma.",
		Mission:     "Impulsar la productividad de las empresas con soluciones tecnológicas confiables, a la medida y con acompañamiento cercano.",
		Vision:      "Ser el aliado tecnológico de referencia para las empresas que buscan crecer con innovación, seguridad y visibilidad total de su operación.",
		Email:       "contacto@ngnnetvision.com",
		Phone:       "+52 55 0000 0000",
		WhatsApp:    "5215500000000",
		Address:     "Ciudad de México, México",
		Hours:       "Lunes a viernes, 9:00 – 18:00",
	},
	Stats: []Stat{
		{Value: "+120", Label: "Proyectos entregados"},
		{Value: "+50", Label: "Clientes activos"},
		{Value: "+800", Label: "Vehículos monitoreados"},
		{Value: "24/7", Label: "Soporte y monitoreo"},
	},
	Services: []Service{
		{
			ID:          "software",
			Icon:        "code",
			Title:       "Desarrollo de Software",
			Subtitle:    "Web, apps y sistemas a la medida",
			Summary:     "Sitios web, aplicaciones móviles y software empresarial diseñados para automatizar y hacer crecer tu negocio.",
			Description: "Convertimos tus procesos en herramientas digitales. Diseñamos, programamos y mantenemos plataformas web, apps móviles y sistemas internos con arquitectura moderna, seguros y listos para escalar.",
			Features: []Feature{
				{
					Title:       "Sitios y aplicaciones web",
					Description: "Sitios corporativos, landing pages y aplicaciones web rápidas, seguras y adaptadas a cualquier pantalla, pensadas para atraer clientes y facilitar tu operación.",
					Time:        "2 a 4 semanas",
					Includes:    []string{"Diseño moderno y adaptable a celular", "Optimización para Google (SEO)", "Panel para editar tu contenido", "Hosting, dominio y certificado SSL", "Formularios de contacto y botón de WhatsApp", "Blog o sección de noticias"},
				},
				{
					Title:       "Apps móviles iOS y Android",
					Description: "Aplicaciones para tus clientes o tu equipo de trabajo, publicadas en App Store y Google Play, con una sola base de código para reducir costos y tiempos.",
					Time:        "2 a 4 meses",
					Includes:    []string{"Diseño de experiencia de usuario (UX/UI)", "Notificaciones push", "Funcionamiento sin conexión", "Publicación en las tiendas de apps", "Inicio de sesión y perfiles de usuario", "Pagos dentro de la app"},
				},
				{
					Title:       "Software a la medida (ERP, CRM, inventarios)",
					Description: "Sistemas hechos exactamente para tus procesos: ventas, clientes, inventarios, compras o producción, en una sola plataforma.",
					Time:        "2 a 6 meses",
					Includes:    []string{"Análisis de tus procesos actuales", "Usuarios, roles y permisos", "Reportes y tableros de indicadores", "Migración de tu información existente", "Control de inventarios y almacenes", "Facturación electrónica (CFDI 4.0)"},
				},
				{
					Title:       "APIs e integraciones con terceros",
					Description: "Conectamos tus sistemas entre sí y con servicios externos para que la información fluya sin capturas dobles.",
					Time:        "1 a 6 semanas",
					Includes:    []string{"Integración con facturación electrónica (CFDI)", "Pasarelas de pago y paqueterías", "Conexión con ERP, CRM y hojas de cálculo", "APIs documentadas y seguras", "Sincronización automática de datos", "Integración con WhatsApp Business"},
				},
				{
					Title:       "Tiendas en línea y pagos digitales",
					Description: "Tu tienda abierta las 24 horas, con catálogo, carrito de compra y cobros con tarjeta, transferencia o pago en tiendas de conveniencia.",
					Time:        "3 a 8 semanas",
					Includes:    []string{"Catálogo con inventario en tiempo real", "Pagos con tarjeta, SPEI y OXXO", "Cálculo de envíos y seguimiento de pedidos", "Cupones, promociones y reportes de ventas", "Conexión con Mercado Libre y Amazon", "Cuentas de cliente e historial de compras"},
				},
				{
					Title:       "Cloud, DevOps y mantenimiento",
					Description: "Mantenemos tus sistemas en línea, actualizados y respaldados, en servidores en la nube que crecen según tu demanda.",
					Time:        "Servicio mensual continuo",
					Includes:    []string{"Servidores en AWS, Azure o Google Cloud", "Respaldos automáticos", "Monitoreo y alertas de disponibilidad", "Actualizaciones de seguridad", "Migración de servidores a la nube", "Soporte técnico con tiempo de respuesta garantizado"},
				},
			},
			Technologies: []string{"React", "Go", "Node.js", "Flutter", "PostgreSQL", "Docker", "AWS"},
			Benefits: []string{
				"Automatiza tareas repetitivas y reduce errores manuales",
				"Información de tu negocio centralizada y en tiempo real",
				"Código propio: el sistema es tuyo y crece contigo",
				"Entregas por etapas para ver avances desde la primera semana",
			},
			IdealFor: []string{"Pymes en crecimiento", "Comercios y tiendas", "Empresas de servicios", "Startups"},
			FAQ: []FAQ{
				{Question: "¿Cuánto tarda un proyecto?", Answer: "Un sitio web toma de 2 a 4 semanas; un sistema a la medida, de 2 a 6 meses según su alcance. Te damos un calendario detallado en la propuesta."},
				{Question: "¿Qué pasa después de la entrega?", Answer: "Ofrecemos garantía, capacitación a tu equipo y planes de mantenimiento con mejoras continuas."},
				{Question: "¿Pueden mejorar un sistema que ya tengo?", Answer: "Sí. Auditamos tu sistema actual y proponemos mejoras, migraciones o integraciones sin detener tu operación."},
			},
		},
		{
			ID:          "redes",
			Icon:        "network",
			Title:       "Redes y Telecomunicaciones",
			Subtitle:    "Infraestructura y conectividad empresarial",
			Summary:     "Diseño, instalación y soporte de redes de datos, fibra óptica, Wi-Fi empresarial y videovigilancia.",
			Description: "Construimos la infraestructura sobre la que corre tu operación. Desde el cableado estructurado hasta la seguridad perimetral, entregamos redes estables, rápidas y monitoreadas.",
			Features: []Feature{
				{
					Title:       "Cableado estructurado certificado",
					Description: "Instalación ordenada y certificada de la red de datos de tu edificio, la base para una conexión estable durante muchos años.",
					Time:        "1 a 3 semanas",
					Includes:    []string{"Cableado categoría 6 y 6A", "Racks, gabinetes y organizadores", "Certificación de cada punto de red", "Etiquetado y planos de la instalación", "Reordenamiento de cableado existente", "Instalación eléctrica regulada y UPS"},
				},
				{
					Title:       "Redes LAN/WAN y Wi-Fi empresarial",
					Description: "Diseñamos redes internas y enlaces entre sucursales, con Wi-Fi de cobertura total y acceso controlado para empleados e invitados.",
					Time:        "1 a 2 semanas",
					Includes:    []string{"Estudio de cobertura Wi-Fi", "Red separada para invitados", "Conexión segura entre sucursales (VPN)", "Administración centralizada de equipos", "Portal cautivo con registro de visitantes", "Wi-Fi para exteriores y naves industriales"},
				},
				{
					Title:       "Instalación y fusión de fibra óptica",
					Description: "Enlaces de fibra óptica de alta velocidad entre edificios, pisos o naves, inmunes a interferencias eléctricas.",
					Time:        "3 días a 3 semanas",
					Includes:    []string{"Tendido aéreo, subterráneo o interior", "Fusión y conectorización", "Pruebas con OTDR y medición de pérdidas", "Reparación de enlaces dañados", "Enlaces entre edificios o naves", "Distribuidores y charolas de fibra"},
				},
				{
					Title:       "Configuración de routers, switches y firewalls",
					Description: "Configuramos y aseguramos tus equipos de red para que funcionen al máximo y estén protegidos contra amenazas.",
					Time:        "2 a 10 días",
					Includes:    []string{"Segmentación de la red (VLANs)", "Firewall y filtrado de contenido", "Balanceo de varios proveedores de internet", "Control de ancho de banda", "Acceso remoto seguro para empleados", "Auditoría de seguridad de la red"},
				},
				{
					Title:       "CCTV y videovigilancia IP",
					Description: "Sistemas de cámaras en alta definición que puedes revisar desde tu celular, con grabación continua y detección inteligente.",
					Time:        "2 a 7 días",
					Includes:    []string{"Cámaras IP de alta definición", "Grabación en NVR o en la nube", "Visualización remota desde el celular", "Detección de movimiento y alertas", "Visión nocturna y cámaras para exterior", "Control de acceso y videoporteros"},
				},
				{
					Title:       "Telefonía IP y enlaces dedicados",
					Description: "Comunicación de voz moderna sobre tu red de datos y conexiones a internet dedicadas con velocidad garantizada.",
					Time:        "1 a 2 semanas",
					Includes:    []string{"Conmutador IP (PBX) y extensiones", "Menú de bienvenida (IVR)", "Extensiones en celular y computadora", "Enlaces simétricos con respaldo", "Grabación de llamadas", "Portabilidad de tus números actuales"},
				},
			},
			Technologies: []string{"Cisco", "MikroTik", "Ubiquiti", "Fortinet", "Fibra óptica", "VoIP"},
			Benefits: []string{
				"Conexión estable y rápida en todas tus áreas",
				"Menos caídas de red y tiempos de inactividad",
				"Protección contra accesos no autorizados",
				"Infraestructura documentada y lista para crecer",
			},
			IdealFor: []string{"Oficinas y corporativos", "Naves industriales", "Escuelas", "Hoteles y comercios"},
			FAQ: []FAQ{
				{Question: "¿Hacen visitas técnicas?", Answer: "Sí. Realizamos un levantamiento en sitio para revisar tu instalación actual y diseñar la red adecuada."},
				{Question: "¿Trabajan con mi equipo actual?", Answer: "Aprovechamos el equipo que esté en buen estado y te recomendamos solo lo necesario para mejorar."},
				{Question: "¿Ofrecen soporte después de la instalación?", Answer: "Contamos con pólizas de soporte y monitoreo remoto para detectar y resolver fallas a tiempo."},
			},
		},
		{
			ID:          "tracking",
			Icon:        "truck",
			Title:       "Rastreo Vehicular GPS",
			Subtitle:    "Control total de tu flotilla",
			Summary:     "Ubicación en tiempo real, historial de rutas, geocercas y alertas para proteger y optimizar tus vehículos.",
			Description: "Instalamos dispositivos GPS y te damos una plataforma web y móvil para saber dónde está cada unidad, cómo se conduce y cuánto consume. Más seguridad, menos costos operativos.",
			Features: []Feature{
				{
					Title:       "Ubicación en tiempo real",
					Description: "Ve en el mapa dónde está cada una de tus unidades en este momento, con actualizaciones cada pocos segundos.",
					Time:        "Instalación en 1 día",
					Includes:    []string{"Actualización cada 10 segundos", "Vista de toda la flotilla en un mapa", "Estado de la unidad: encendida, detenida o en movimiento", "Compartir ubicación con tus clientes", "Vista de calle y tráfico en tiempo real", "Rastreo de remolques y cajas secas"},
				},
				{
					Title:       "Historial de rutas y reportes",
					Description: "Consulta los recorridos de cualquier día y recibe reportes automáticos de la actividad de tu flotilla.",
					Time:        "Disponible desde el primer día",
					Includes:    []string{"Reproducción de recorridos", "Kilometraje y tiempo de manejo", "Paradas y tiempos de espera", "Reportes en PDF y Excel por correo", "Historial guardado hasta 12 meses", "Reportes por operador y por unidad"},
				},
				{
					Title:       "Geocercas y alertas instantáneas",
					Description: "Define zonas en el mapa y recibe un aviso cuando una unidad entra, sale o se desvía de su ruta.",
					Time:        "Configuración en 1 día",
					Includes:    []string{"Zonas ilimitadas en el mapa", "Alertas por app, SMS y correo", "Aviso de desvío de ruta", "Alertas fuera de horario laboral", "Aviso de llegada a clientes", "Alerta por desconexión del equipo GPS"},
				},
				{
					Title:       "Paro de motor remoto",
					Description: "Detén una unidad a distancia en caso de robo, de forma segura y solo con autorización.",
					Time:        "Instalación en 2 a 3 horas",
					Includes:    []string{"Activación desde la app o la plataforma", "Bloqueo seguro a baja velocidad", "Botón de pánico para el operador", "Apoyo en la recuperación de la unidad", "Sensor de apertura de puertas", "Micrófono para escucha remota en cabina"},
				},
				{
					Title:       "Control de velocidad y combustible",
					Description: "Vigila cómo se manejan tus vehículos y cuánto combustible consumen para reducir gastos y riesgos.",
					Time:        "Instalación en 1 día",
					Includes:    []string{"Alertas por exceso de velocidad", "Detección de frenadas y aceleraciones bruscas", "Sensor de nivel de combustible", "Aviso de cargas o descargas sospechosas", "Calificación de manejo por operador", "Control de tiempo en ralentí (motor encendido detenido)"},
				},
				{
					Title:       "Plataforma web y app móvil",
					Description: "Controla toda tu flotilla desde la computadora o el celular, con acceso para todo tu equipo.",
					Time:        "Acceso inmediato",
					Includes:    []string{"Acceso las 24 horas desde cualquier lugar", "Usuarios con distintos permisos", "Recordatorios de mantenimiento", "Integración con tus sistemas vía API", "Capacitación para tu equipo", "Soporte técnico 24/7"},
				},
			},
			Technologies: []string{"GPS/GNSS", "4G LTE", "Plataforma web", "App móvil", "Alertas SMS/Email", "API"},
			Benefits: []string{
				"Reduce robos y facilita la recuperación de unidades",
				"Ahorro en combustible al optimizar rutas",
				"Mejores hábitos de manejo de tus operadores",
				"Reportes automáticos para tomar mejores decisiones",
			},
			IdealFor: []string{"Flotillas de reparto", "Transporte de carga", "Empresas de servicio en campo", "Vehículos particulares"},
			FAQ: []FAQ{
				{Question: "¿Cuánto tarda la instalación?", Answer: "Entre 1 y 2 horas por vehículo. Podemos instalar en tus instalaciones para no afectar tu operación."},
				{Question: "¿Puedo ver mis unidades desde el celular?", Answer: "Sí. Tienes acceso a la plataforma web y a la app móvil las 24 horas."},
				{Question: "¿Funciona en cualquier tipo de vehículo?", Answer: "Sí: autos, camionetas, camiones, motocicletas y maquinaria."},
			},
		},
	},
	Highlights: []Highlight{
		{Icon: "layers", Title: "Soluciones integrales", Description: "Software, redes y rastreo con un solo proveedor y un solo punto de contacto."},
		{Icon: "users", Title: "Equipo especializado", Description: "Ingenieros y desarrolladores con experiencia en proyectos empresariales."},
		{Icon: "shield", Title: "Seguridad primero", Description: "Buenas prácticas de ciberseguridad en cada sistema, red y dispositivo."},
		{Icon: "clock", Title: "Soporte 24/7", Description: "Monitoreo continuo y atención rápida cuando más lo necesitas."},
	},
	Process: []Step{
		{Icon: "search", Title: "Diagnóstico", Description: "Escuchamos tus necesidades y analizamos tu operación actual."},
		{Icon: "file", Title: "Propuesta", Description: "Diseñamos una solución a la medida con alcance, tiempos y costos claros."},
		{Icon: "zap", Title: "Implementación", Description: "Desarrollamos e instalamos con entregas visibles y pruebas de calidad."},
		{Icon: "lifebuoy", Title: "Soporte continuo", Description: "Te acompañamos con mantenimiento, monitoreo y mejoras constantes."},
	},
}
