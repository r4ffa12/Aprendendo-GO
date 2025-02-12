// Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
// Na prática:
// %d %b
// x << y
// iota * 10 << 10 = kb, mb, gb

package main

import (
	"fmt"
)

func main() {

	y := 24

	x := y << 2

	z := y >> 2

	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)

}
