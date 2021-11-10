# Gestores de tareas y dependecias

## Gestor de tareas

Como task runner se ha seleccionado [Task](https://taskfile.dev/#/).

### ¿Qué es?

Es una herramienta de ejecución / contrucción de tareas que pretende ser más simple y fácil de usar que, por ejemplo, GNU Make.

Es una buena elección para el proyecto ya que Task está escrito en Go. No tiene otras dependencias, lo que significa que no necesita configuraciones de instalación complejas.

### Características

* Instalación sencilla, se tiene que agregar un binario y agregarlo al $PATH.

* Multiplataforma, también disponible para Windows y macOS.

* Excelente para la generación de código, se puede evitar facilmente que una tarea se ejecute si un conjunto determinado de archivos no ha cambiado desde la última ejecución.

### Uso

Task tiene una documentación muy amplia y correcta, para verla podemos hacerlo desde [aquí](https://taskfile.dev/#/usage). Para comenzar a usar Task se tiene que crear un fichero `Taskfile.yml` en la carpeta raiz de nuestro proyecto, a este fichero se incorporará las lineas de texto necesarias para hacer referencia a la carpeta donde se encuentran los ficheros seguido de los comandos suficientes para realizar la acción que se desea.

## Gestor de dependecias

Cuando el código usa paquetes externos, esos paquetes (distribuidos como módulos) se convierten en dependencias. Con el tiempo, es posible que deba actualizarlos o reemplazarlos. Go proporciona herramientas de gestión de dependencias que le ayudan a mantener seguras sus aplicaciones Go a medida que incorpora dependencias externas.

Toda la documentación acerca del gestor de dependecias se puede encontrar en la página principal de [golang](https://golang.org/doc/modules/managing-dependencies).
