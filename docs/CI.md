# Integración Continua

Se trata de implementar acciones que se ejecuten cuando sucede un evento en nuestro repositorio, en general, se deberán de pasar nuestros tests desde una máquina externa, quien ejecutará estos tests y nos dirá cuales han pasado y cuales no.

## Elección de servicio online de CI

Para completar el objetivo se debe de implementar al menos dos servicios de integración continua. Primero vemos requisitos a tener en cuenta para seleccionar un servicio.

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

### Motivos

En el caso de **Circle CI**, este sistema es compatible con imagenes Docker, por lo que podemos aprovechar que tenemos una imagen de nuestro proyecto en Docker Hub y usarla para ejecutar los tests. Si no quisieramos usar nuestra imagen podemos optar por usar Matrix jobs, que era uno de los requisitos que buscabamos de un sistema de integración continua, en nuestro caso no lo vamos a usar pero aun así [aquí](https://circleci.com/blog/circleci-matrix-jobs/) se encuentra la documentación sobre como montar. En este caso Circle CI tiene completa compatiblidad con Github útil para los Checks API y por último tiene una configuración muy sencilla, iniciando sesión en la plataforma con los datos de github, gracias al OAuth que ofrece, directamente detecta nuestro repositorios e indicar cual es el que queremos configurar. Al ser ficheros de configuración .yml es muy similar la sintaxis que vamos a usar respecto a la que hemos usado hasta ahora, por ejemplo la usada para configurar workflows de Github. Al final queda un fichero muy compacto, con pocas directivas y que funciona a la perfección. Otra característica es que ofrece plantillas para diversos lenguajes, en nuestro caso no va a servir para mucho pero es algo que cabe mencionar para tener en cuenta en alguna otra ocasión si se usa un lenguaje distinto al usado en este proyecto.

En el caso de **Github Actions**, me optado por este sistema por su plena compatibilidad con Github, como es obvio. Ofrece Matrix jobs, los cuales vamos a configurar y nos puede resultar útil que podamos instalar nuestro task runner para poder usar las ordenes que tengamos configuradas en nuestro repositorio. Por último, la configuración es muy similar a la que utilizamos para publicar una imagen de nuestro proyecto en Docker Hub por tanto estaba totalmente familiarizado con a sintaxis usada y entendiendo todas las directivas desde el primer momento que se monta el fichero, solo se tiene que buscar como usar matrix jobs.

Otra posible elección podía ser **Travis CI**, ofrece todo lo mencionado que buscamos en sistemas CI pero al intentar probar el producto tenía que insertar la tarjeta de crédito, algo que me echó bastante para atras para probar el sistema.

