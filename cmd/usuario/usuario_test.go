package usuario

import (
	"fmt"
	"testing"
)

func TestObjetivo3_usuario(t *testing.T){
	nombre := "Luis"
	apellido := "Arostegui Ruiz"
	fnac := "26/04/1999"
	usuario := NewUsuario(nombre, apellido, fnac)
	fmt.Println(usuario)
}