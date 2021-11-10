package MoneyAccount

import "time"

type mipin interface {
}

type OpcionAccount func(*Account)

func Nombre(nombre string) OpcionAccount {
	return func(acco *Account) {
		acco.Nombre = nombre
	}
}

func Saldo(saldo float64) OpcionAccount {
	return func(acco *Account) {
		acco.saldo = saldo
	}
}

func Fcreate(fcreate time.Time) OpcionAccount {
	return func(acco *Account) {
		acco.fcreate = fcreate
	}
}

func Descripcion(des string) OpcionAccount {
	return func(acco *Account) {
		acco.Descripcion = des
	}
}
