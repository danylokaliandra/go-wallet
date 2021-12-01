package usuario

type Usuario struct {
	nombre   string
	apellido string
	fnac     string
}

func NewUsuario(name string, apellido string, fnac string) *Usuario {
	user := new(Usuario)
	user.nombre = name
	user.apellido = apellido
	user.fnac = fnac
	return user
}
