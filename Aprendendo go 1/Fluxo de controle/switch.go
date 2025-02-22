/*package main

import (
	"fmt"
)

func main() {

	x := 10

	switch true {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")

	}
}
*/

/*
package main

import (
	"fmt"
)

func main() {

	quemestanoescritoriohoje := "zezinho"

	switch quemestanoescritoriohoje {

	case "zezinho":
		fmt.Println("hoje quem ta no escritorio é o zezinho")
		fallthrough
	case "marquinho":
		fmt.Println("sempre que o zezinho vem, o marquinho vem tambem")
	case "joana":
		fmt.Println("hoje quem tá no escritorio é a joana")
		fallthrough
	case "maria":
		fmt.Println("hoje quem tá no escritorio é a maria")
	}

}
hoje quem ta no escritorio é o zezinho
sempre que o zezinho vem, o marquinho vem tambem */















package main

import (
	"fmt"
)

func main() {

	quemestanoescritoriohoje := "ninguem"

	switch quemestanoescritoriohoje {

	case "zezinho":
		fmt.Println("hoje quem ta no escritorio é o zezinho")
		fallthrough
	case "marquinho":
		fmt.Println("sempre que o zezinho vem, o marquinho vem tambem")
	case "joana":
		fmt.Println("hoje quem tá no escritorio é a joana")
		fallthrough
	case "maria":
		fmt.Println("hoje quem tá no escritorio é a maria")
	default:
		fmt.Println("hoje é feriado")
	}

}
hoje é feriado 

