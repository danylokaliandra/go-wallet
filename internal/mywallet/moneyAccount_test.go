package mywallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjetivoAhorroMensual(t *testing.T) {

	assert := assert.New((t))
	t.Log("Test calcular objetivo ahorro mensual")

	cuenta := NewAccount("Viaje", 2000.50, 40000)
	cuenta.AniadirBalance(2000)
	cuenta.AniadirBalance(-50)
	cuenta.AniadirBalance(-19.99)
	cuenta.AniadirBalance(-5)
	var anios int = 5

	objetivoMensual := cuenta.ObjetivoAhorroMensual(anios)
	wantedvalue := 633.325

	assert.Equal(objetivoMensual, wantedvalue, "Valor obtenido erroneo")
}
