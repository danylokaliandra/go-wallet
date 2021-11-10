package Usuario

import "time"

type Usuario struct {
	Nombre   string
	Apellido string
	fnac     time.Time
	sexo     typeSexo
}

func NewUsuario(opciones ...Opcion) Usuario {
	user := Usuario{Nombre: "N/A", Apellido: "N/A",
		fnac: time.Date(1997, time.January, 1, 0, 0, 0, 0, time.Local),
		sexo: MASCULINO}
	for _, opcion := range opciones {
		opcion(&user)
	}
	return user
}

func (user Usuario) GetEdad() uint8 {
	return uint8(time.Now().Year() - user.fnac.Year())
}
