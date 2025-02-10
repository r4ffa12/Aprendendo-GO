package main

import (
	"fmt"
)

func main() {
	a := "e"
	b := "é"
	c := "香"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v", d, e, f)

	s := "ascii éøâ 香"

	for _, v := range s {
		fmt.Printf("\n%b - %T - %#U - %#x\n", v, v, v, v)

	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
