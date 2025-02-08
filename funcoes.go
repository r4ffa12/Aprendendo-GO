package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)

	sub := subtracao(10, 5)
	fmt.Println(sub)

	nome, sobrenome := printaNomeCompleto("Rafa", "Braulio")
	fmt.Println(nome)
	fmt.Println(sobrenome)

}

// Função começando com letra minúscula:
// Função é PRIVADA
// Função ela só pode ser utilizada com próprio pacote
func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra Maiuscula
// Função é Pública
// Função ela pode ser ultilizada em outros pacotes
func PrintaNome(nome string) string {
	return nome
}
func soma(x int, y int) int {

	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}
