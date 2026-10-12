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
	Benefits      []string `json:"benefits"`
	BenefitsTitle string   `json:"benefitsTitle"` // opcional; por defecto "Beneficios para tu empresa"
	IdealFor      []string `json:"idealFor"`
	FAQ           []FAQ    `json:"faq"`
	// Proveedores que se muestran en la sección "Nuestros proveedores".
	Providers []Provider `json:"providers"`
}

type Provider struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
			Providers: []Provider{
				{Name: "AWS", Description: "Servidores y servicios en la nube"},
				{Name: "Docker", Description: "Contenedores para desplegar aplicaciones"},
				{Name: "PostgreSQL", Description: "Bases de datos confiables"},
				{Name: "React", Description: "Interfaces web modernas"},
				{Name: "Go", Description: "Backends rápidos y seguros"},
				{Name: "Node.js", Description: "APIs y servicios web"},
				{Name: "Flutter", Description: "Apps para iOS y Android"},
			},
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
			Description: "Diseñamos, instalamos y mantenemos la infraestructura de red de tu negocio para que tu equipo trabaje sin interrupciones, de forma rápida y segura.",
			Features: []Feature{
				{
					Title:       "Redes empresariales",
					Description: "Instalamos y configuramos tu red desde cero o reorganizamos la que ya tienes.",
					Includes: []string{
						"Cableado estructurado Cat6, racks y etiquetado profesional",
						"Configuración de switches y routers",
						"Segmentación de red (VLANs) para oficina, invitados, cámaras y telefonía",
					},
				},
				{
					Title:       "Wi-Fi empresarial",
					Description: "Cobertura estable en toda tu oficina, restaurante, hotel o bodega.",
					Includes: []string{
						"Estudio de cobertura",
						"Instalación de access points con conexión continua al moverte",
						"Red de invitados separada y portal de acceso personalizado",
					},
				},
				{
					Title:       "Seguridad de red",
					Description: "Protegemos tu información y tu conexión.",
					Includes: []string{
						"Firewall y filtrado de contenido",
						"VPN para trabajo remoto",
						"Conexión segura entre sucursales",
					},
				},
				{
					Title:       "Internet sin interrupciones",
					Description: "Tu negocio siempre conectado.",
					Includes: []string{
						"Doble proveedor de internet con respaldo automático",
						"Respaldo con 4G/5G",
						"Balanceo de carga entre enlaces",
					},
				},
				{
					Title:       "Enlaces inalámbricos",
					Description: "Conectamos edificios, naves o bodegas sin necesidad de cableado entre ellos.",
				},
				{
					Title: "Telefonía IP",
					Includes: []string{
						"Conmutadores IP, extensiones y menú de bienvenida (IVR)",
						"Llamadas claras y estables con calidad de servicio configurada",
					},
				},
				{
					Title: "Videovigilancia IP",
					Includes: []string{
						"Instalación de cámaras IP con red dedicada",
						"Acceso remoto seguro desde tu celular",
					},
				},
				{
					Title:       "Diagnóstico de red",
					Description: "¿Tu internet está lento o se cae? Revisamos tu red, encontramos el problema y te entregamos un reporte con soluciones.",
				},
				{
					Title:       "Soporte y mantenimiento mensual",
					Description: "Olvídate de los problemas de red con nuestras pólizas de soporte:",
					Includes: []string{
						"Monitoreo de tu red",
						"Soporte remoto",
						"Visitas programadas y atención a fallas",
					},
				},
			},
			Technologies: []string{"Cisco", "MikroTik", "Ubiquiti", "Fortinet", "Fibra óptica", "VoIP"},
			Providers: []Provider{
				{Name: "Cisco", Description: "Switches, routers y seguridad empresarial"},
				{Name: "MikroTik", Description: "Routers y enrutamiento"},
				{Name: "Ubiquiti", Description: "Wi-Fi y enlaces inalámbricos"},
				{Name: "Fortinet", Description: "Firewalls y seguridad de red"},
				{Name: "Palo Alto Networks", Description: "Firewalls de nueva generación y ciberseguridad"},
				{Name: "Meraki", Description: "Redes administradas desde la nube"},
			},
			BenefitsTitle: "¿Por qué elegirnos?",
			Benefits: []string{
				"Personal certificado en Cisco (CCNA)",
				"Soluciones a la medida de tu negocio y presupuesto",
				"Instalaciones limpias, ordenadas y documentadas",
				"Atención directa y rápida",
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
			Description: "Seguimiento, supervisión y gestión remota en tiempo real de vehículos, contenedores, activos y personas. Administra todas tus unidades desde una sola plataforma web y app móvil, conectado con tu empresa 24/7.",
			Features: []Feature{
				{
					Title:       "Mapa en tiempo real",
					Description: "Consulta tus unidades, geocercas, recorridos y eventos sobre un mapa detallado que muestra el estado y la condición de cada una en tiempo real.",
					Includes: []string{
						"Ubicación y estado de cada unidad",
						"Geocercas y recorridos en el mapa",
						"Marcadores de eventos",
						"Detección de la ubicación del dispositivo",
					},
				},
				{
					Title:       "Lista de trabajo",
					Description: "Controla en tiempo real los parámetros de movimiento, encendido y ubicación de tus unidades, y qué tan actualizados están sus datos.",
					Includes: []string{
						"Movimiento de las unidades",
						"Encendido y apagado (ignición)",
						"Ubicación actual",
						"Actualidad de los datos",
					},
				},
				{
					Title:       "Geocercas",
					Description: "Crea zonas de acuerdo con tus necesidades y ve en qué geocerca está cada unidad en lugar de su dirección.",
					Includes: []string{
						"Geocercas personalizadas",
						"Ubicación por geocerca o por dirección",
					},
				},
				{
					Title:       "Historial de recorridos",
					Description: "Revisa en orden cronológico los eventos de cada unidad y visualízalos en el mapa.",
					Includes: []string{
						"Viajes",
						"Estacionamientos",
						"Llenados y descargas de combustible",
					},
				},
				{
					Title:       "Informes y estadísticas",
					Description: "Informes operativos y estadísticos, con gráficas detalladas, para analizar patrones y tomar mejores decisiones.",
					Includes: []string{
						"Viajes y paradas",
						"Llenados y descargas de combustible",
						"Comportamiento del conductor y del vehículo",
						"Horas de trabajo",
					},
				},
				{
					Title:       "Comandos remotos",
					Description: "Gestiona tus unidades a distancia con comandos totalmente personalizados.",
					Includes: []string{
						"Envío de mensajes y rutas",
						"Solicitud de configuración",
						"Solicitud de instantáneas",
					},
				},
				{
					Title:       "Alertas y eventos personalizados",
					Description: "Crea eventos predefinidos y recibe alertas en tiempo real desde cualquier lugar y en cualquier momento.",
					Includes: []string{
						"Exceso de velocidad en un área específica",
						"Apertura de la puerta de carga fuera de la zona autorizada",
						"Notificaciones a la medida de tu operación",
						"Parámetros de seguridad y mejora continua",
					},
				},
				{
					Title:       "Plataforma web y app móvil",
					Description: "Plataforma web y app para Android y iOS: monitorea y controla tu operación las 24 horas, los 365 días del año, desde cualquier computadora o celular.",
					Includes: []string{
						"Alertas en tiempo real",
						"Ubicación de tus dispositivos",
						"Detalles de cada dispositivo",
						"Lista de notificaciones con su detalle y ubicación",
					},
				},
				{
					Title:       "Contenedores, carga y activos",
					Description: "Seguimiento y gestión de contenedores, mercancías, maquinaria y otros activos.",
					Includes: []string{
						"GPS para contenedores",
						"Seguimiento y gestión de mercancías",
						"Monitoreo satelital",
						"Candado inteligente que reporta",
					},
				},
				{
					Title:       "Localizador personal",
					Description: "Seguimiento en tiempo real de tu personal para su protección.",
				},
			},
			Technologies: []string{"GPS/GNSS", "Monitoreo satelital", "Plataforma web", "App Android e iOS", "Candado inteligente"},
			// Proveedores de GPS por definir: mientras esté vacío se muestra "Próximamente".
			Providers: []Provider{},
			Benefits: []string{
				"Conectado con tu empresa 24/7, los 365 días del año",
				"Todas tus unidades en una sola plataforma",
				"Información en tiempo real en todos tus dispositivos",
				"Datos para tomar mejores decisiones y mejorar la seguridad",
			},
			IdealFor: []string{"Flotillas de reparto", "Transporte de carga y mercancías", "Contenedores y maquinaria", "Protección de personal"},
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
