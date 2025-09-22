package main

import "fmt"

func main() {

	// Primeira maneira
	func() {
		fmt.Println("Olá mundo!")
	}() // Assim que foi declarada, aqui já é chamada.... Executado no caso.

	// Outra maneira
	func(texto string) {
		fmt.Println(texto)

	}("Olá mundo 2")

	// Outra maneira

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d %.2f", texto, 10, 12.00000)

	}("Passando parâmetro")
	fmt.Println(retorno)
}
