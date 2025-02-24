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

/*
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

*/

/*
package main

import (
	"fmt"

)



func main() {

var x interface{}

	x = true

	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
	}
}

é um bool
*/

/*
package main

import (
	"fmt"
)

var x interface{}

func main() {

	switch x := 2; {
	case x == 1:
		fmt.Println("é um 1")
	case x == 2:
		fmt.Println("é um 2")
	case x == 3:
		fmt.Println("é um 3")
	case x == 4:
		fmt.Println("é um 4")

	}

}
	é um 2
*/