# Tests Unitarios

Para realizar el objetivo 4 de la asignatura se pide documentar posibles opciones de bibliotecas de asercciones y justificar la elección de alguna de las opciones.

## Frameworks de tests para Golang

Go es un lenguaje sin excepciones. El fin de las expeciones es que cuando se produce una excepción, el programa no debería de continuar sino pararse. En go tenemos una función, *panic*, que se usa cuando queramos que el programa se pare. Aun asi no deberíamos de usar *panic* muy a menudo. Otra opción que tenemos es hacer un control de errores. Como en go se pueden devolver más de un valor en una función, se puede devolver el valor esperado y un error. Esto podemos personalizarlo según como nos sea más comodo y haya un mayor control. Por ejemplo, podemos crear nuestra propia estructura con un mensaje y código de error, creamos funciones para que crear el error y mostrar el error obtenido. Esto conlleva declarar condicionales cada vez que queramos controlar algo de nuestro código, en mi opinión algo laborioso.

Para ejecutar test en go se puede hacer uso de una biblioteca de testing, nos encontramos con distitntas opciones:

* **Testify**, amplía el testing ligero de Go para realizar afirmaciones y simulacros de dependencias. Permite generar datos automáticamente para los tests.
* **gocheck**, posee informes de errores útiles para ayudar a resolver problemas.
* **gopwt**, muestra de una manera más gráfica el error en un test. Permite usar operadores binarios para la comprobación de la ejeccución de un test.
* **go-testdeep**, permite probar fácilmente las rutas HTTP API, utilizando operadores flexibles y el ayudante tdhttp.
* **Ginkgo**, se basa en el paquete de pruebas de Go, lo que permite pruebas de estilo expresivo de desarrollo impulsado por el comportamiento ("BDD"). Por lo general, se empareja con la biblioteca de comparadores de Gomega. Framework bastante pesado.
* **Goblin**, permite probar cosas asincrónicas, como goroutines, etc.
* **GoConvey**, otro framework de testeo BDD.

## Elección y motivación

En el proyecto se usará [Testify](https://github.com/stretchr/testify). Aunque se puede usar el propio framework de tests con el paquete *testing* considero que este es bastante ligero y el control de errores durante la ejecución de test es bastante más laboriosa que si se utilizase Testify. Posee una buena documentación y es recomendado en la documentación oficial de Golang. Es de los más populares para este lenguaje de programación. Posee un gran número de métodos para aserciones y permite crear testing suite para realizar el setup y teardown si los test son más complejos. La integración con Go es clara, se permite ejecutar *go test* para ejecutar todos los test que tengamos. Apoyandons en task, nuestro task runner seleccionado en el anterior objetivo, nuestro proyecto se vuelve potente y sencillo a la hora de ejecutar los test.

## Uso de la herramienta seleccionada

Se ha creado un fichero llamado 'moneyAccount_test.go' en el que se declaran los test correspondientes a la entidad Usuario y moneyAccount.