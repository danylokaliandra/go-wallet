# MyWallet :moneybag:

## Idea Inicial :bulb:

La idea es crear un sistema que permita a los usuarios tener el control sobre sus gastos. La aplicación mostrará un informe actualizado dia a dia sobre los gastos/ingresos del usuario, esto lo podrá hacer automatizando gastos/ingresos mensuales para asi predecir el dinero que va a poder gastar en un periodo de tiempo conseguiendo al mismo tiempo ahorrar dinero. Se calcularán los gastos de los usuarios para asi reducir los gastos de su dia a dia para obtener unos ahorros más consolidados.

## Motivación :high_brightness:

La idea del proyecto viene de querer predecir el dinero que generarán los usuarios en un tiempo determinado en base a sus gastos e ingresos para lograr una salud financiera y evitar problemas con su dinero. El principal problema que se trata de solucionar es el de ahorrar diaria, semanal, mensual o anualmente según como quiera el usuario. La idea es que se ajuste al objetivo de ahorros del usuario.

## Documentación :bookmark_tabs:

Para acceder a toda la documentación puede acceder desde [aquí](docs/)

## Task runner para el proyecto

Como este proyecto se esta implementando el Go, se ha buscado un task runner para este. El elegido es [Task](https://taskfile.dev/#/). Para ver más información sobre el gestor de tareas y dependecias pinche [aquí](/docs/gestor.md).

1. Necesitamos tener instalado tener instalado en nuestro sistema **Go**. Si no esta instalado se puede hacer desde [aquí](https://golang.org/doc/install).

2. Para instalar Task podemos hacerlo desde un script, esto resulta util en escenarios como CI. La siguiente orden instalará el gestor de tareas en el directorio /usr/local/bin:

	```shell

	sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

	```


3. Clonamos el repositorio y accedemos a el:

	``` shell
	
	git clone git@github.com:LuisArostegui/MyWallet.git
	cd MyWallet
	
	```

4. Uso de Task
Una vez en el repositorio podemos ejecutar las siguientes acciones:
	* **task install**, se generará un ejecutable del proyecto en $GOPATH/bin
	* **task run**, ejecuta el proyecto.
	* **task check**, se comprobará la sintaxis de las entidades programadas.
	* **task installdeps**, se incluirán las dependecias del proyecto en el fichero go.mod.
	* **task test**, se ejecutarán todos los tests del proyecto.
	 


## Test unitarios

Se puede consultar toda la información acerca de los test unitarios desarrollados [aquí](docs/tests.md)