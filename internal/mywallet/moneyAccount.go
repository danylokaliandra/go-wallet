package mywallet

type Account struct {
	nombre string
	saldo  float64
}

func NewAccount(nombre string, saldo float64) *Account {
	acc := new(Account)
	acc.nombre = nombre
	acc.saldo = saldo
	return acc
}

func (a *Account) GetSaldo() float64 {
	return a.saldo
}

func (a *Account) SetSaldo(saldo float64) {
	a.saldo = saldo
}

func (a *Account) GetNombre() string {
	return a.nombre
}

func (a *Account) SetNombre(nombre string) {
	a.nombre = nombre
}
