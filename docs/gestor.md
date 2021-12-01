# Gestor de tareas y dependecias

## Gestor de tareas

En Golang nos podemos encontrar los siguientes task runner:

* Realize:
	* Podemos manejar varios proyecto simultaneamente.
	* Acepta todos los comandos de Go.
	* Más info: github.com/oxequa/realize
* Godo:
	* Es un task runner y analizador de ficheros de Golang basado en Rake, gulp.
	* Más info: https://github.com/go-godo/godo
* Myke:
	* Similar a make pero con yaml.
	* myke resuelve todos estos problemas en un único binario diminuto, para evitar reinventar las mismas cosas una y otra vez.
	* Más info: https://github.com/omio-labs/myke
	
Se usará Task como task runner del proyecto.

### Motivación

Go es un lenguaje que usa un task runner implícito, con subcomandos de go se puede compilar (go build), ejecutar tests (go test) o instalar un paquete (go install) pero tiene sus limites y puede ser necesario añadir herramientas externas para cubrir estas limitaciones. 

Algunas de esas limitaciones las podemos encontrar a la hora de ejecutar varios tests, si solo tenemos un test bastaría con ejecutar "go test *ruta-del-test*", pero si tenemos varios test puede ser molesto especificar rutas concretas del proyecto para que se ejecute todo correctamente. 

Herramientas que nos ayuden a cubrir estas necesidades pueden ser Task o make. Ambas herramientas necesitan de un fichero donde se especifiquen las acciones que queremos realizar en nuestro proyecto, este fichero se puede ir ampliando conforme vayan aparenciendo nuevas necesidades en nuestra aplicación.

Cualquiera de los dos gestores de tareas para la asignatura son suficientes, el hecho de seleccionar Task en lugar de make es:

- La documentación y uso de la herramienta. Por ejemplo, el poder indicar y enviar variables a nuestro proyecto, útil para indicar en un futuro varios tiempos de datos a nuestra aplicación como en esta se van a escribir números donde pueden salir numeros más grandes, más pequeños, con más o menos decimales es importante que sea sencillo poder probar varios tipos de números y escribirlo continuamente en una terminal va a ser molesto por esto podemos crear indicando un fichero de texto por separado indicando las variables y desde el Taskfile.yml usar estas variables para pasarlas como parámetros a distintas funciones de nuestro proyecto.
- Podemos incluir otros taskfiles, esto puede ser útil para un futuro para cuando despleguemos la aplicación usando docker.

### Características de Task

* Instalación sencilla, se tiene que agregar un binario y agregarlo al $PATH.

* Multiplataforma, también disponible para Windows y macOS.

* Excelente para la generación de código, se puede evitar facilmente que una tarea se ejecute si un conjunto determinado de archivos no ha cambiado desde la última ejecución.

### Uso

Task tiene una documentación muy amplia y correcta, para verla podemos hacerlo desde [aquí](https://taskfile.dev/#/usage). Para comenzar a usar Task se tiene que crear un fichero `Taskfile.yml` en la carpeta raiz de nuestro proyecto, a este fichero se incorporará las lineas de texto necesarias para hacer referencia a la carpeta donde se encuentran los ficheros seguido de los comandos suficientes para realizar la acción que se desea.

## Gestor de dependecias

Cuando el código usa paquetes externos, esos paquetes (distribuidos como módulos) se convierten en dependencias. Con el tiempo, es posible que deba actualizarlos o reemplazarlos. Go proporciona herramientas de gestión de dependencias que le ayudan a mantener seguras sus aplicaciones Go a medida que incorpora dependencias externas.

Toda la documentación acerca del gestor de dependecias se puede encontrar en la página principal de [golang](https://golang.org/doc/modules/managing-dependencies).
