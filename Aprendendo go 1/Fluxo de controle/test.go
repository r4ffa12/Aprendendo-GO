/*
package main

import (

	"fmt"

)

func main() {

		if x := 101; x > 100 {
			fmt.Println("x é maior que 100")
		} else if x < 10 {
			fmt.Println("x é menor que 10")
		} else {
			fmt.Println("x não é maior que 10 nem maior que 100")
		}
	}
*/
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Valor de i:", i)
	}
}
