# Contenedor de pruebas

# Motivación

Los contenedores son una forma optimizada de crear, probar y poner en marcha aplicaciones en varios entornos. Algunos de sus beneficios son:

* Menos sobrecarga, requieren menos recursos del sistema que los enternos de máquinas virtuales o hardware.
* Mayor portabilidad, las aplicaciones se pueden poner en marcha fácilmente en sistemas operativos y platadormas de hardware diferentes.
* Mayor eficiencia, permiten escalar las aplicaciones con mayor rapidez.
* Mejor desarrollo de aplicaciones, los contenedores respaldan los esfuerzos ágiles y de DevOps para acelerar los ciclos de desarrollo, prueba y producción.

En nuestro proyecto vamos a usar [Docker](https://www.docker.com/). Aparte de Docker existen otras herramientas que pueden realizar la misma función de esta, estonces, ¿porqué usar Docker?

Algunas de las cosas que buscamos en un contenedor son:

* Las instancias deben de ser ligeras, con un contenedor se tiene que conseguir un montaje sencillo de nuestra aplicación haciendo que este se arranque rápidamente.
* Deben de ser consistentes, cuando se desarrolla en un contenedor se hacen pruebas de la aplicación dentro de este y se despliegan dentro de un contenedor. Esto significa que el entorno de pruebas es idéntico al entorno en el que se va a ejecutar el software.
* Eficiencia de imágenes del contenedor, poder crear una imagen de contenedor y usar esa misma imagen a lo largo de todo el proceso de despliegue.

Docker cumple con estos requisitos además de que:

* Es gratuito y de código abierto
* Ofrece compatibilidad y mantenimiento más facil.
* Las configuraciones son rápidas y simples
* Garantiza entornos consistentes desde la fase de desarrollo hasta la fase de producción, es decir, ofrecen pruebas continuas.

## Alternativas a Docker

Docker no es perfecto, se puede mencionar que una de sus desventajas es que debe de ejecutarse con privilegios root y al parar el contenedor se elimina toda la información dentro de él a excepción de lo que está en los volúmenes. Otra desventaja es que no divide los recursos sino que consume los recursos del host a demanda, esto puede generar que otros contenedores vayan lento por culpa de otro que consuma demasiados recursos. Por tanto otras opciones a Docker son:

* RKT, uno de los mayores competidores de Docker. Su fuerte es la seguridad. Antes de la versión 1.1 Docker necesitaba correrse como usuario root, una vulnerabilidad grave que permite ataques a nivel de super usuario. RKT permite utilizar el manejo estandar de grupos para permisos Linux, esto permite ejecutar el contenedor una vez creado con un usuario sin privilegios root. Docker tiene la ventaja de que es más fácil de integrar mientras que RKT conlleva una instalación y configuración más manual.
* Podman, nos permite ejecutar contenedores con usuarios sin privilegios root. Soporta múltiples formatos de imágenes como OCI y Docker. Soporta el manejo de Pods.
* LXC, usa una tecnología de virtualización a nivel de sistema operativo que permite crear y correr múltiples entornos virtuales linux de manera aislada.

Vistas estas alternativas vamos a ver como integrar Docker en nuestro proyecto.

## Integración de Docker en nuestro proyecto

### Imagen de Golang

Docker tiene muchas imagenes de golang, cada una de ellas diseñada para un caso de uso específico:

* `golang:<version>`, es la imagen por defecto. Si no estamos seguros de cuáles son nuestras necesidades, probablemente esta es la mejor opción. Además puede incluir etiquetas como pueden ser *bullseye*, *buster* o *stretch*. Estas etiquetas son los nombres de código de la suite para las versiones de Debian e indican en que versión se basa la imagen. 
* `golang:<version>-alpine`, esta imagen se basa en el proyecto Alpine Linux. Las imagenes Alpine Linux son mucho más livianas que la mayoría de imágenes base de distribución (~5 MB). Esta variante es experimental y no es oficialmente compatible con el [proyecto Go](https://github.com/golang/go/issues/19938). La principal advertencia a tener en cuenta es que utiliza **musl libc** en lugar de **glibc**, puede llegar a provocar un comportamiento inesperador en nuestra aplicación. En [este artículo](https://news.ycombinator.com/item?id=10782897) se conversa acerca de los problemas que puede traer este tipos de imagenes.
* `golang:<version>-windowsservercore`, esta imagen se basa en Windows Server Core.

En nuestro caso tenemos que debatir se seleccionar la imagen basada en Debian o en Alpine. Como hemos comentado la principal diferencia entre estas es el tamaño y Alpine viene con la desventaja de que la imagen es una variante experimental pero esto para nuestro proyecto en un principio no acarrea ningún problema y se va a optar por la imagen Alpine por su menor tamaño respecto a Debian. En concreo la versión de la imagan va a ser la 1.17, hay una versión 1.18 a dia de hoy, 28/12/2021, pero es una versión beta.

### Facilitar uso de Docker con nuestro task runner

Se ha automatizado la ejecución de los tests en el task runner. Esto se consigue añadiendo la siguiente directiva a nuestro Taskfile.yml:

```docker run -t -v `pwd`:/app/test luisarostegui/mywallet```

### Buenas prácticas en nuestro Dockerfile

* Vamos a usar una imagen ligera (Alpine) para optimizar el tamaño y poder tener el control sobre los paquetes necesarios.
* Usar variables con ENV para directorios de trabajo.
* Ejecutar tanto las instalaciones de las dependencias como el task runner como usuario y no como superusuario.

### Justificación de directivas en el Dockerfile

1. Con la directiva FROM especificamos la imagen base.
2. La directiva LABEL la utilizamos para especificar nombre y correo de la persona encargada del Docker.
3. Con ENV especificamos la ruta donde queremos que se ejecuten los test, sirve para crear una variable de entorno.
4. RUN lo usamos para ejecutar ordenes de terminal, tales como crear un grupo y un usuario.
5. USER lo usamos para cambiar de usuario.
6. La directiva COPY permite copiar los ficheros de dependecias a la carpeta /app/test.
7. WORKDIR, especificamos la ruta donde queremos trabajar.
8. ENTRYPOINT, indicamos la acción a ejecutar, en este caso `task test`.