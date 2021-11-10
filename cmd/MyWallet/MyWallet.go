package MyWallet

import . "MyWallet/pkg/Usuario"

func main() {

	user := NewUsuario(Nombre("hola"))
	Nombre("hola")
	println(user.Nombre)
}
