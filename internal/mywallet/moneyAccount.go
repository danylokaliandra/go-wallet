package mywallet

type Account struct {
	nombre  string
	total   float64
	balance map[string]float64
}

func NewAccount(nombre string) *Account {
	acc := new(Account)
	acc.nombre = nombre
	acc.balance = make(map[string]float64)
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

func (a *Account) AniadirBalance(nombre string, valor float64) {
	a.balance[nombre] = valor
}

func (a *Account) GetBalance() map[string]float64 {
	return a.balance
}
