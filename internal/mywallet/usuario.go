package mywallet

import (
	"time"
)

type Usuario struct {
	nombre  string
	fnac    time.Time
	cuentas []Account
}

func NewUsuario(name string, fnac time.Time) *Usuario {
	user := new(Usuario)
	user.nombre = name
	user.fnac = fnac
	return user
}

func (u *Usuario) GetFnac() time.Time {
	return u.fnac
}

func (u *Usuario) SetFnac(fnac time.Time) {
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

func (u *Usuario) PredecirAhorrosEdad(edad int) float64 {
	año_usuario, _, _ := u.fnac.Date()
	fecha_actual := time.Now()
	año_actual, _, _ := fecha_actual.Date()
	edad_actual := int(año_actual - año_usuario)
	edad_ahorrando := edad - edad_actual
	ahorro := 0.0

	for _, cuenta := range u.cuentas {
		ahorro_anual := cuenta.PredecirAhorrosAnuales()
		ahorro += (ahorro_anual * float64(edad_ahorrando))
	}

	return ahorro
}
