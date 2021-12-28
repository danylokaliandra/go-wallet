#Imagen base para docker
FROM golang:1.17-alpine

# Metadatos de información del encargado de mantenimiento
LABEL maintainer="Luis Aróstegui Ruiz <luisarostegui@correo.ugr.es>"

# Creamos variable de entorno para el directorio donde vamos a ejecutar los tests
ENV TEST_DIR=/app/test

# Añadimos usuario sin privilegios de superusuario y cremos un grupo para dicho usuario
RUN addgroup -S mywallet && adduser -S mywallet -G mywallet

# Cambiamos al nuevo usuario
USER mywallet

#Establecemos el directorio donde vamos a ejecutar los tests con nuestro nuevo usuario
WORKDIR $TEST_DIR

#Instalamos modulos necesarios para compilar
COPY go.mod ./
COPY go.sum ./

#Ahora podemos descargar y actualizar las dependecias
RUN go mod download

#Instalamos nuestro task runner
RUN go install github.com/go-task/task/v3/cmd/task@latest

#Especificamos el ejecutable que usará el contenedor
ENTRYPOINT ["task", "test"]