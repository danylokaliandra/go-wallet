package mywallet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObjetivoAhorroMensual(t *testing.T) {

	assert := assert.New((t))
	t.Log("Test calcular objetivo ahorro mensual")

	total := 2000.50
	objetivo := 40000.5
	assert.Greater(objetivo, total, "Error. Total es mayor que el objetivo")
	cuenta := NewAccount("Viaje", total, objetivo)
	cuenta.AniadirBalance(2000)
	cuenta.AniadirBalance(-450.66)
	var anios int = 5

	objetivoMensual := cuenta.ObjetivoAhorroMensual(anios)
	wantedvalue := 633.3333333333334

	assert.Equal(objetivoMensual, wantedvalue, "Valor obtenido erroneo")
}

func TestPredecirAhorrosAnuales(t *testing.T) {

	assert := assert.New((t))
	t.Log("Test predecir ahorros anuales")

	cuenta := NewAccount("Banco", 10000.00, 30000.65)
	cuenta.AniadirBalance(3000)
	cuenta.AniadirBalance(-50)
	cuenta.AniadirBalance(-400)

	ahorroAnual := cuenta.PredecirAhorrosAnuales()
	wantedvalue := 40600.0

	assert.Equal(ahorroAnual, wantedvalue, "Valor obtenido erroneo")
}

func TestAsignarCuentas(t *testing.T) {
	assert := assert.New((t))
	t.Log("Test crear y asignar cuentas a un usuario")

	tTime := time.Date(1999, time.April, 26, 0, 0, 0, 0, time.Local)
	usuario := NewUsuario("Luis", tTime)

	cuenta := NewAccount("Banco", 10000.00, 30000.65)
	cuenta.AniadirBalance(3000)

	total := 2000.50
	objetivo := 40000.5
	assert.Greater(objetivo, total, "Error. Total es mayor que el objetivo")
	cuenta_2 := NewAccount("Viaje", total, objetivo)
	cuenta_2.AniadirBalance(2000)

	usuario.AniadirCuenta(*cuenta)
	usuario.AniadirCuenta(*cuenta_2)
	wantedlen := 2

	assert.Equal(len(usuario.GetCuentas()), wantedlen, "No se han añadido correctamente las cuentas")
}
