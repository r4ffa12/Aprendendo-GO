// Iota¶
// Em uma declaração constante, o identificador pré-declarado representa constantes inteiras não tipadas sucessivas. Seu valor é o índice do respectivo ConstSpec nessa declaração constante, começando em zero. Ele pode ser usado para construir um conjunto de constantes relacionadas: iota

package main

import (
	"fmt"
)

// Numa declaração de constantes, o identificador iota representa números sequenciais.
const (
	a = iota * 10 //pode se colocar so uma vez e fazer fomrula
	_             // iota
	c             // iota
	x             // iota
	_             // iota
	z             // iota
)

func main() {
	fmt.Println(a, c, x, z)
	// print = 0 2 3 5
	// print2 = 10 20 30 50
}
