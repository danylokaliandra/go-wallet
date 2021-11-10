package MoneyAccount

import (
	"time"
)

type cuenta interface {
	GetSaldo() float64
	GetidAccount() string
}

type Account struct {
	Nombre      string
	saldo       float64
	fcreate     time.Time
	Descripcion string
}

type BankAccount struct {
	Account
	idAccount string
}

func NewAccount(opciones ...OpcionAccount) Account {
	acco := Account{saldo: 0,
		fcreate: time.Date(1997, time.January, 1, 0, 0, 0, 0, time.Local)}
	for _, opcion := range opciones {
		opcion(&acco)
	}
	return acco
}

func NewBankAccount(id string, opciones ...OpcionAccount) BankAccount {
	acco := BankAccount{Account: NewAccount(opciones...), idAccount: id}

	return acco
}

func (a Account) GetSaldo() float64 {
	return a.saldo
}

func (a *Account) SetSaldo(saldo float64) {
	Saldo(saldo)(a)
}

func (a Account) GetFcreate() time.Time {
	return a.fcreate
}

func (a *Account) SetFcreate(fcreate time.Time) {
	Fcreate(fcreate)(a)
}

func (ba BankAccount) GetId() string {
	return ba.idAccount
}

func (ba BankAccount) SetId(id string) {
	ba.idAccount = id
}

/* func (a Account) GetTiempoCreacion() string{
	return uint8(time.Now().Year() - a.fnac.Year())
} */
