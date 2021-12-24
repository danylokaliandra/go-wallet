package mywallet

type Account struct {
	nombre   string
	total    float64
	objetivo float64
	balance  []float64
}

func NewAccount(nombre string, total float64, objetivo float64) *Account {
	acc := new(Account)
	acc.nombre = nombre
	acc.total = total
	acc.objetivo = objetivo
	return acc
}

func (a *Account) GetTotal() float64 {
	return a.total
}

func (a *Account) SetTotal(total float64) {
	a.total = total
}

func (a *Account) GetObjetivo() float64 {
	return a.objetivo
}

func (a *Account) SetObjetivo(objetivo float64) {
	a.objetivo = objetivo
}

func (a *Account) GetNombre() string {
	return a.nombre
}

func (a *Account) SetNombre(nombre string) {
	a.nombre = nombre
}

func (a *Account) AniadirBalance(valor float64) {
	a.balance = append(a.balance, valor)
}

func (a *Account) GetBalance() []float64 {
	return a.balance
}

func (a *Account) ObjetivoAhorroMensual(anios int) float64 {
	var ahorro_mensual float64 = 0

	ahorro_mensual = (a.objetivo - a.total) / (float64(anios) * 12)

	return ahorro_mensual
}

func (a *Account) PredecirAhorrosAnuales() float64 {
	var ahorro_anual float64 = 0
	var balance_mensual float64 = 0

	for _, valor := range a.balance {
		balance_mensual += valor
	}

	ahorro_anual = (balance_mensual * 12) + a.total

	return ahorro_anual
}
