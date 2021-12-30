package mywallet

import (
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestObjetivoAhorroMensual(t *testing.T) {
	g := Goblin(t)

	total := 2000.50
	objetivo := 40000.5
	cuenta := NewAccount("Viaje", total, objetivo)
	cuenta.AniadirBalance(2000)
	cuenta.AniadirBalance(-450.66)
	var anios int = 5

	objetivoMensual := cuenta.ObjetivoAhorroMensual(anios)
	wantedValue := 633.3333333333334

	g.Describe("Objetivo Ahorro Mensual", func() {
		// Passing Test
		g.It("objetivoMensual debe coincidir con wantedValue", func() {
			g.Assert(objetivoMensual).Equal(wantedValue)
		})
	})
}

func TestPredecirAhorrosAnuales(t *testing.T) {
	g := Goblin(t)

	cuenta := NewAccount("Banco", 10000.00, 30000.65)
	cuenta.AniadirBalance(3000)
	cuenta.AniadirBalance(-50)
	cuenta.AniadirBalance(-400)

	ahorroAnual := cuenta.PredecirAhorrosAnuales()
	wantedValue := 40600.0

	g.Describe("Objetivo Ahorro Mensual", func() {
		// Passing Test
		g.It("ahorroAnual debe coincidir con wantedValue", func() {
			g.Assert(ahorroAnual).Equal(wantedValue)
		})
	})
}

func TestAsignarCuentas(t *testing.T) {
	g := Goblin(t)

	tTime := time.Date(1999, time.April, 26, 0, 0, 0, 0, time.Local)
	usuario := NewUsuario("Luis", tTime)

	cuenta := NewAccount("Banco", 10000.00, 30000.65)
	cuenta.AniadirBalance(3000)

	total := 2000.50
	objetivo := 40000.5
	cuenta_2 := NewAccount("Viaje", total, objetivo)
	cuenta_2.AniadirBalance(2000)

	usuario.AniadirCuenta(*cuenta)
	usuario.AniadirCuenta(*cuenta_2)
	wantedlen := 2

	g.Describe("Objetivo Ahorro Mensual", func() {
		// Passing Test
		g.It("Longitud del slice debe coincidir con en numero de cuentas añadidas al usuario", func() {
			g.Assert(len(usuario.GetCuentas())).Equal(wantedlen)
		})
	})
}

func TestPredecirAhorrosEdad(t *testing.T) {
	g := Goblin(t)

	tTime := time.Date(1999, time.April, 26, 0, 0, 0, 0, time.Local)
	usuario := NewUsuario("Luis", tTime)

	cuenta := NewAccount("Banco", 600.0, 30000.65)
	cuenta.AniadirBalance(200)

	total := 500.0
	objetivo := 40000.5
	cuenta_2 := NewAccount("Viaje", total, objetivo)
	cuenta_2.AniadirBalance(300)

	usuario.AniadirCuenta(*cuenta)
	usuario.AniadirCuenta(*cuenta_2)

	edad_ahorrando := 55
	ahorroEdad := usuario.PredecirAhorrosEdad(edad_ahorrando)
	wantedValue := 234300.0

	g.Describe("Objetivo Ahorro Mensual", func() {
		// Passing Test
		g.It("ahorroEdad debe coincidir con wantedValue", func() {
			g.Assert(ahorroEdad).Equal(wantedValue)
		})
	})
}
