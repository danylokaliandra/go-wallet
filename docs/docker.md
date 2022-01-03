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

### Cuestiones a tener en cuenta en la elección de una imagen

Nos vamos a encontrar, en general, con dos tipos de imagen distintas, una basada en debian y otra en Alpine. Las diferencias entre estas dos son las siguientes:

* El tamaño entre una y otra. Alpine tiene un tamaño menor que Debian. Con esto conseguimos que operaciones como construir, pull y push sean más rápidas.
* Consumir menos memoria por el propio sistema operativo en comparación con Debian. Alpine se considera seguro y rápido.

### Imagen de Golang

Docker tiene muchas imagenes de golang, cada una de ellas diseñada para un caso de uso específico:

* `golang:<version>`, es la imagen por defecto. Si no estamos seguros de cuáles son nuestras necesidades, probablemente esta es la mejor opción. Además puede incluir etiquetas como pueden ser *bullseye*, *buster* o *stretch*. Estas etiquetas son los nombres de código de la suite para las versiones de Debian e indican en que versión se basa la imagen. 
* `golang:<version>-alpine`, esta imagen se basa en el proyecto Alpine Linux. Las imagenes Alpine Linux son mucho más livianas que la mayoría de imágenes base de distribución (~5 MB). Esta variante es experimental y no es oficialmente compatible con el [proyecto Go](https://github.com/golang/go/issues/19938). La principal advertencia a tener en cuenta es que utiliza **musl libc** en lugar de **glibc**, puede llegar a provocar un comportamiento inesperador en nuestra aplicación. En [este artículo](https://news.ycombinator.com/item?id=10782897) se conversa acerca de los problemas que puede traer este tipos de imagenes.
* `golang:<version>-windowsservercore`, esta imagen se basa en Windows Server Core.

En nuestro caso tenemos que debatir se seleccionar la imagen basada en Debian o en Alpine. Como hemos comentado la principal diferencia entre estas es el tamaño y Alpine viene con la desventaja de que la imagen es una variante experimental pero esto para nuestro proyecto en un principio no acarrea ningún problema y se va a optar por la imagen Alpine por su menor tamaño respecto a Debian. En concreo la versión de la imagen va a ser la 1.17, hay una versión 1.18 a dia de hoy, 28/12/2021, pero es una versión beta, la version 1.17 es la última más estable actualmente.

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

## Uso de Docker Hub

Para realizar la configuración de mi repositorio con Docker Hub he consultado [esta página](https://docs.docker.com/docker-hub/builds/).

Se tiene que crear un Github Action que crea la imagen del contenedor y la publica en Docker Hub de forma automática. Con esto conseguiremos que cada vez que avancemos de objetivo (se haga un push a la rama principal), se automatice la construcción de la imagen. También me ha sido de utilidad la siguiente [documentación](https://docs.docker.com/ci-cd/github-actions/).

### Construcción de nuestro fichero para el workflow

1. Indicamos cuando se debe de publicar la imagen en docker hub.

```yaml
on:
  push:
    branches: # Indicamos la rama de nuestro repositorio que queremos analizar.
      - main
    paths: # Indicamos los ficheros que tiene que analizar para realizar la publicación de la imagen.
      - '**.go' #  Si estos ficheros no se han modificado no se realiza la publicación
      - go.mod
  pull_request:
    branches:
      - main
```
* Bajo mi punto de vista, si se realiza un push a la rama principal y hemos modificado algún fichero .go o el go.mod deberíamos de actualizar la imagen ya que habremos añadido código que afectará a la lógica de negocio de nuestro proyecto. Otra opción puede ser añadir `path-ignore` e incluir los ficheros, por ejemplo, de documentación, con esto conseguiriamos que si se edita un fichero de documentación y se hace push a la rama principal no se actualice la imagen en docker hub.
* Cuando hacemos merge desde un pull request a la rama principal, en la asignatura significa que siempre se van a realizar cambios respecto a la lógica de negocio de nuestro proyecto, es decir, vamos a estar realizando cambios importantes en nuestro proyecto por esto pienso que para este repositorio concreto si hacemos merge de un pull request a la rama main se ejecute la publicación de la imagen.

2. Creo una variable para especificar el repositorio del que queremos crear y publicar una imagen.
```yaml
env:
  REPO: mywallet
```
3. Especificamos que queremos que suceda dentro de nuestro flujo de trabajo.
```yaml
jobs:
  build:
    runs-on: ubuntu-latest # Indicamos que se ejecute en las últimas instancias de Ubuntu disponibles.
    
    steps:
      - name: Checkout # Revisa nuestro repo en $GITHUB_WORKSPACE para que nuestro workflow pueda acceder a el.
        uses: actions/checkout@v2
      - name: Login to Docker Hub # Iniciamos sesión en docker hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}
      - name: Build and push # Contruimos la imagen y la publicamos
        uses: docker/build-push-action@v2
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKER_HUB_USERNAME }}/$REPO:latest
```


# Cambios de framework de testing

Tras completar el fichero de CI workflow, en nuestro caso [este](.github/workflows/dockerhub.yml), me han surgido problemas para poder ejecutar el Dockerfile. Aparentemente, el Dockerfile tenía buena pinta no había error sintáticos. Pero al ejecutar ```docker run -t -v `pwd`:/app/test luisarostegui/mywallet``` surgió el siguiente error:

```console
 Objetivo-5 U:5 ?:1  ~/UGR/IV/Mi-repo/MyWallet                                                                                                                      12:54:09  luismsi 
❯ task docker
task: [docker] docker run -t -v `pwd`:/app/test luisarostegui/mywallet
task: [test] go test -v ./...
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go: updates to go.mod needed; to update it:
        go mod tidy
task: Failed to run task "test": exit status 1
task: Failed to run task "docker": exit status 1
```

Lo primero que se me vino a la mente fue que hay un error al instalar nuestro task runner y no consegue ejecutar correctamente la sentencia `Task test` correctamente. Ya que el error nos indica que hay paquetes que no están actualizados caí en que el error tenía que estar referido a los paquetes que importo en mi proyecto pero esto no debería de ser un gran problema, `go mod tidy` actualiza nuestras dependecias y `go mod download` las descarga... ¿Entonces porque no consigue descargar las dependecias? El único import que tenía que podía suponer problemas era el de **Testify**, probé a quitar este paquete y a ejecutar los tests sin framework de test y correcto, ahi estaba el error, la versión 1.17 de Go no puede incluir esta dependencia. Como el error parecía de la imagen base seleccionada, opté por seleccionar otra imagen, como por ejemplo una imagen que no fuese Alpine (obtenia el mismo error) u otra versión de Go. La opción de otra versión de Go parecía atractiva ya que pude observar en la documentación de Testify que su proyecto funciona de manera estable en versiones de Go desde la 1.13 hasta la 1.15, es decir, realmente no soporta la versión 1.17 que se estaba usando hasta el momento. Cambiando a una versión 1.15, al ejecutar nuestro docker run obtenía el siguiente error:

```console
  Objetivo-5 U:5 ?:1  ~/UGR/IV/Mi-repo/MyWallet                                                                                                                                13:14:07  luismsi 
❯ task docker
task: [docker] docker run -t -v `pwd`:/app/test luisarostegui/mywallet
task: [test] go test -v ./...
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
FAIL    github.com/LuisArostegui/MyWallet/internal/mywallet [build failed]
FAIL
task: Failed to run task "test": exit status 2
task: Failed to run task "docker": exit status 1
```

Otro error a nuestro catalogo. Al buscar el error parece que tenía que incluir en el Dockerfile una operación como `RUN apk add g++` ya que el error parece que viene de un paquete de Testify que necesita funciones escritas en C. Pero aun incluyendo tenía el mismo error una y otra vez.

Al seguir indagando sobre el error, quizas este estaba en como se montaba el fichero para el workflow de dockerhub, encontré una sección en la documentación oficial para crear este fichero especificamente para este lenguaje, Go, pero al cambiar este fichero no suponía ningún cambio en el error.

Por esto he decidido abandonar Testify y volver a la documentación del anterior objetivo en busca de un framework de test que se ajuste al proyecto. Probé con go-testdeep pero me pareció demasiasdo complejo para los tests que tengo actualmente en mi proyecto. He probado con **Goblin**, del que investigué acerca de este framework en el anterior objetivo y pienso que se adapta perfectamente al proyecto, sencillo crear tests y se informa de manera muy gráfica y agradable si se pasan o no los tests.
