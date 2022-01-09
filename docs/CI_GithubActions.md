# Explicación de Github Actions

## Documentación
Se ha obtenido documentación de [aquí](https://docs.github.com/es/actions/learn-github-actions/managing-complex-workflows#using-a-build-matrix).
Para usar nuestro task runner en el workflow se ha obtenido la documentación de [aquí](https://github.com/arduino/setup-task)

## Contrucción  .github/workflows/tests.yml
Se busca crear un flujo de trabajo donde podamos ejecutar los tests de nuestro proyecto en diferentes versiones de Go. Es importante que estas versiones tengan soporte y sean estables para ello encontré [esta página](https://endoflife.date/go) donde se puede apreciar que las versiones que tienen soporte actualmente son la 1.16 y la 1.17.



``` yaml
name: Tests Github Actions

on: push # Cada vez que se haga un push en el repo se ejecutará el workflow

jobs: # Indicamos que queremos que pase en nuestro workflow
  tests-github: # Nombre del jobs
    strategy: # keyword para montar nuestra matrix
      matrix:
        golang: ['1.16', '1.17'] # Indicamos mediente un array las versiones de Go con las que queremos ejecutar nuestros tests
    name: Run tests Go ${{ matrix.golang }}
    runs-on: ubuntu-latest
    steps:
      - name: Instalar Task # Instalamos nuestro task runner. 
        uses: arduino/setup-task@v1

      - name: Checkout
        uses: actions/checkout@v2

      - name: Setup golang
        uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.golang }}

      - name: Run tests
        run: task test

```

Destacar que instalar nuestro task runner es prescindible, podemos ignonar la instalación y simplemente cuando vayamos a ejecutar los tests indicarlo con `go test -v ./...`.