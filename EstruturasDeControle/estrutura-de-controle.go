package main

import "fmt"

func main() {

	fmt.Println("Estruturas de Controle")

	numero := -10

	if numero > 15 {
		fmt.Println("Numero é maior que 15")
	} else {
		fmt.Println("Numero é menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre 0 e - 10")
	}
}
