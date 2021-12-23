# Tests Unitarios

Para realizar el objetivo 4 de la asignatura se pide documentar posibles opciones de bibliotecas de asercciones y justificar la elección de alguna de las opciones.

## Frameworks de tests para Golang

### ¿Que se busca en framework para tests?

En este caso para un lenguaje de programación como Go, se busca un framework que:

* Se pueda ejecutar con `go test` con el fin de aprovechar el core que Go nos proporciona, como son los paquetes de testing como su binario para la ejecución.
* Sencillez, es decir, que se puedan añadir aserciones facilmente y de manera intutiva.
* Mocking, el proyecto puede llegar a tener un gran número de clases, si por cada test que creamos tenemos que instanciar las clases va a ser un trabajo donde vamos a perder mucho tiempo. Por ello si el framework posee esta habilidad generaremos tests de una manera más eficiente.
* Debe de ser ligero, para que los tests sean suficientemente rápidos como para cumplir los principios de F.I.R.S.T.
* Disponer de un test suite para facilitar el setup y teardown.

Go es un lenguaje sin excepciones. El fin de las expeciones es que cuando se produce una excepción, el programa no debería de continuar sino pararse. En go tenemos una función, *panic*, que se usa cuando queramos que el programa se pare. Aun asi no deberíamos de usar *panic* muy a menudo. Otra opción que tenemos es hacer un control de errores. Como en go se pueden devolver más de un valor en una función, se puede devolver el valor esperado y un error. Esto podemos personalizarlo según como nos sea más comodo y haya un mayor control. Por ejemplo, podemos crear nuestra propia estructura con un mensaje y código de error, creamos funciones para que crear el error y mostrar el error obtenido. Esto conlleva declarar condicionales cada vez que queramos controlar algo de nuestro código, en mi opinión algo laborioso.

Para ejecutar test en go se puede hacer uso de una biblioteca de testing, nos encontramos con distintas opciones:

* **Testify**, amplía el testing ligero de Go para realizar afirmaciones y simulacros de dependencias. Permite generar datos automáticamente para los tests.
* **gocheck**, posee informes de errores útiles para ayudar a resolver problemas.
* **gopwt**, muestra de una manera más gráfica el error en un test. Permite usar operadores binarios para la comprobación de la ejeccución de un test.
* **go-testdeep**, permite probar fácilmente las rutas HTTP API, utilizando operadores flexibles y el ayudante tdhttp.
* **Ginkgo**, se basa en el paquete de pruebas de Go, lo que permite pruebas de estilo expresivo de desarrollo impulsado por el comportamiento ("BDD"). Por lo general, se empareja con la biblioteca de comparadores de Gomega. Framework bastante pesado.
* **Goblin**, permite probar cosas asincrónicas, como goroutines, etc.
* **GoConvey**, otro framework de testeo BDD.

## Elección y motivación

En el proyecto se usará [Testify](https://github.com/stretchr/testify). Aunque se puede usar el propio framework de tests con el paquete *testing* considero que este es bastante ligero y el control de errores durante la ejecución de test es bastante más laboriosa que si se utilizase Testify. Posee una buena documentación y es recomendado en la documentación oficial de Golang. Es de los más populares para este lenguaje de programación. Posee un gran número de métodos para aserciones y permite crear testing suite para realizar el setup y teardown si los test son más complejos. La integración con Go es clara, se permite ejecutar *go test* para ejecutar todos los test que tengamos. Apoyandons en task, nuestro task runner seleccionado en el anterior objetivo, nuestro proyecto se vuelve potente y sencillo a la hora de ejecutar los test.

### Comparación de Testify frente a otros frameworks

* **Testify vs GoConvey**, en este caso tenemos que GoConvey es un framework para tests unitarios y tests de regresión, mientras que Testify es solo un framework para test unitarios, en este caso para el objetivo no se pide nada más que tests unitarios. GoConvey es una herramienta de tests que, por un lado observa el código en busca de cambios y se puede ejecutar mediante 'go test' y muestra sus resultados en un navegador, por otro lado posee una biblioteca que le permite escribir tests estilo BDD, mientras que Testify realiza las mismas funciones que GoConvey salvo el poseer una biblioteca para tests BDD, en este caso Testify, bajo mi punto de vista, las aserciones son más sencillas de escribir, permite mocking (lo que simplifica en gran medida la creación de objetos para ejecutar tests) algo que GoConvey permite pero con una biblioteca de terceros, GoMock. Testify es una herramienta potente y más sencilla que GoConvey que puede llegar a ser más potente que Testify pero es más pesada.

* **Testify vs GoCheck**, esta comparación es algo más compleja. Son muy parecidos, al igual que pasaba con GoConvey, GoCheck no permite mocking y hay que recurrir a GoMock. GoCheck permite fixtures mientras que Testify no lo permite, sin embargo Testify tiene un paquete 'suite' para desarolladores que permite agrupar los tests y construir una tests suite donde implementa una fase de teardown y esta se puede configurar, gocheck no implementa un paquete para tests suite pero si se pueden crear. Cualquiera de las dos opciones es aceptable. Si nos fijamos en las fases de test: *setup*, *tests*, *teardown*, quizas sea más recomendable Testify.

* **Testify vs GoCheck**, muy similares. Goblin es simple y flexible basado en NodeBDD y ofrece una manera de mostrar los reports y sintaxis más amable para el usuario. Pero Goblin no permite Mocks ni test suite. Pasa lo mismo con gopwt.



## Uso de la herramienta seleccionada

Se ha creado un fichero llamado 'moneyAccount_test.go' en el que se declaran los test correspondientes a la entidad Usuario y moneyAccount.

## F.I.R.S.T

Los tests desarrollados siguen el principio F.I.R.S.T:

* Fast, los test deben de ser rápidos, en mi caso los test tardan un tiempo promedio de 0.0001s después de haber ejecutado 10 veces los tests.
* Isolated, cada test ejecuta una parte independiente de la entidad.
* Repeteable, siempre se obtiene el mismo resultado intependientemente de la máquina donde se ejecute.
* Self-validating, los test, gracias a Testify, se comprueban automáticamente si se pasan. No hace falta validarlo manualmente.
* Thorough, los test cubren todos los escenarios posibles.