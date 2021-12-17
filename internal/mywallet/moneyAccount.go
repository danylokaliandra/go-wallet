package mywallet

type Account struct {
	nombre string
	total  float64
}

func NewAccount(nombre string) *Account {
	acc := new(Account)
	acc.nombre = nombre
	return acc
}

func (a *Account) GetTotal() float64 {
	return a.total
}

func (a *Account) SetTotal(total float64) {
	a.total = total
}

func (a *Account) GetNombre() string {
	return a.nombre
}

func (a *Account) SetNombre(nombre string) {
	a.nombre = nombre
}
