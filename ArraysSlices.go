package main

import "fmt"

func main() {
	// Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	// fmt.Println(array[0])				//Hello
	// fmt.Println(array[1])				//World
	// fmt.Println(array[0], array[1])		//Hello World
	// fmt.Println(array)   				// [Hello World]

	numPrimo := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimo)
	fmt.Println(numPrimo[0:3])
	fmt.Println(numPrimo[1:])

	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	slice[2] = "Rafael"
	fmt.Println(slice[2])
	fmt.Println(slice)

	//fmt.Println(len(slice)) vai mostrar quantidades de slice nesse ex: 5

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14)
	fmt.Println(numPares)
}

// Listas

//
// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo
// [1, 2, 3, 3, 4] - [] int
// ["Rafael", "Santos", "programação"] - []string
// [1.353, 2.424, 3.55] - [] float64

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chaves - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro tipo
// map [string] int
// { "rafael": 25, "lucia": 55}
// map [string]string
// {"rafael": "Santos", "lucia": Santos}

// ARRAY
// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho de array
// Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// SLICE

// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[a], a[1]
// Função embutida len() retorna o tamanho do slice
// Função append() usada para adicionar valores
