# Integración Continua

Se trata de implementar acciones que se ejecuten cuando sucede un evento en nuestro repositorio, en general, se deberán de pasar nuestros tests desde una máquina externa, quien ejecutará estos tests y nos dirá cuales han pasado y cuales no.

## Elección de servicio online de CI

Para completar el objetivo se debe de implementar al menos dos servicios de integración continua. Primero vemos requisitos a tener en cuanta para seleccionar un servicio.

### Que buscamos en un servicio de CI

* **Configuración sencilla**, la herramienta CI debería permitir una fácil instalación y configuración.
* **Facilidad de uso**, debería de funcionar completamente en segundo plano y tiene que ser sencilla su construcción y ver los resultados de la implementación.
* **Arquitectura extensible**, que tenga complementos nos permitirá agregar funcionalidades a nuestra herramienta CI.
* **Gratuito**.
* **Matrix jobs**, con esto conseguimos ejecutar tests en varias versiones que nosotros indiquemos.

### Algunos servicios de CI

* **Jenkins**, servidor de automatización de código abierto para la integración continua, es bastante popular y de código abierto. El servidor está disponible para Windows, Mac y Linux. Tiene una configuración sencilla y agradable para el usuario gracias a su GUI. Diseñado como un servidor de automatización extensible, esto quiere decir que se puede usar solo para un servidor de CI o se puede convertir en un centro de entrega continua completo. Ofrece notificaciones sobre el estado de compilación. Posee Matrix Jobs.
* **Bamboo**, destaca por mostrar métricas de calidad y estado. Una función de Agent Matrix permite asignar compilaciones a los agentes y permite visualizar los requisitos del sistema para cada compilación. Bamboo se encarga de las tareas de detección, compilación, prueba y fusión de ramas para implementar el código de dorma continua en entornos de producción o de prueba, basándose en el nombre de la rama.
* **CircleCI**, compatible con muchos lenguajes de programación. Su única opción para el control de versiones es Github. Su GUI es bastante agradable para el usuario, proporciona paneles de control con estadísticas y datos sobre como trabajan los equipos y como se ejecuta su código. No es una herramienta gratuita pero tiene un nivel gratuito.
* **Buddy**, destaca por su tiempo promedio de implementación, de unos 12 segundos, o por su procedimiento de configuración . Usa pipelines para construir, probar e implementar software. Ofrece una prueba de 5 proyectos gratis, a partir de que este se supere se tendrá que pagar por su uso.
* **GoCD**, servidor de CI de código abierto, que se utiliza para visualizar y modelar fácilmente flujos de trabajo complejos. Tiene una interfaz intuitiva para crear canalizaciones de CD. Las canalizaciones se pueden tratar como un código normal registrado en el control de código fuente, lo que permite el control de versiones y el seguimiento de las canalizaciones. Adminte formatos JSON y YAML para la configuración.
* **Github Actions**, es sencillo su uso. Tiene integración total, evidentemente, con Github. Posee Matrix Jobs.
* **GitLab**, su herramienta CI se incluye como una aplicación web con una API abierta que administra proyectos a través de una interfaz de usuario amigable, integrándose con todas las funciones de GitLab. Ofrece un plan gratuito con todas las etapas del ciclo de vida de DevOps y hasta 2000 minutos de CI / CD. Posee Matrix Jobs.

## Elección

Se elegirá Circle CI y Github Actions como sistemas de CI para el proyecto.


