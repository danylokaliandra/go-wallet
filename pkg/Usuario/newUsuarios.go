package Usuario

import (
	"time"
)

type typeSexo uint8

const (
	MASCULINO typeSexo = iota // 0
	FEMENINO                  // 1
)

func (s typeSexo) String() string {
	switch s {
	case 0:
		return "Masculono"
	case 1:
		return "Femenino"
	}
	return "Invalido"
}

type Opcion func(*Usuario)

func Nombre(nombre string) Opcion {
	return func(user *Usuario) {
		user.Nombre = nombre
	}
}

func Apellido(apellido string) Opcion {
	return func(user *Usuario) {
		user.Apellido = apellido
	}
}

func Fnac(fnac time.Time) Opcion {
	return func(user *Usuario) {
		user.fnac = fnac
	}
}

func Sexo(sexo typeSexo) Opcion {
	return func(user *Usuario) {
		user.sexo = sexo
	}
}
