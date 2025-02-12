package main

import (
	"fmt"
)

// Constantes
// São valores imutáveis
// const oi = "Bom dia"
// const oi string = "Bom dia"
// As não tipadas só terão um tipo atribuido a elas quando forem usadas.
// Ex. qual o tipo de 42? int? uint? float64?
// Ou seja, é uma flexibilidade conveniente.
// Na prática: int, float, string.
// const x = y
// const ( x = y ) -  para declarar varias

const (
	x = 10
	y = 20
	z = 30
)

// var y = 10

// var c int

var d float64

func main() {

	fmt.Println(x, y, z)
}
