package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if 1%2 != 0 {
			fmt.Println(i)
		}
	}
}