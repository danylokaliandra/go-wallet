package mywallet

type Usuario struct {
	nombre  string
	fnac    string
	cuentas []Account
}

func NewUsuario(name string, fnac string) *Usuario {
	user := new(Usuario)
	user.nombre = name
	user.fnac = fnac
	return user
}

func (u *Usuario) GetFnac() string {
	return u.fnac
}

func (u *Usuario) SetFnac(fnac string) {
	u.fnac = fnac
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) GetCuentas() []Account {
	return u.cuentas
}

func (u *Usuario) AniadirCuenta(cuenta Account) {
	u.cuentas = append(u.cuentas, cuenta)
}
