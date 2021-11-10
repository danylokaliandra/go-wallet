package moneyAccount

import (
	"fmt"
	"testing"
)

func TestObjetivo3_cuenta(t *testing.T){
	nombre_cuenta := "Efectivo"
	saldo := 99.99
	cuenta := NewAccount(nombre_cuenta, saldo)
	fmt.Println(cuenta)
}