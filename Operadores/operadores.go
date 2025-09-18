package main

import "fmt"

func main() {

	// ARITIMÉTICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM ARITIMÉTICOS


	// ATRIBUIÇÃO

	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM ATRIBUIÇÃO


	// OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM DOS OPERADORES RELACIONAIS


	//OPERADORES LÓGICOS
	fmt.Println("----------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println(verdadeiro && !falso)

	// FIM DOS OPERADORES LÓGICOS

	// OPERADOES UNÁRIOS
	fmt.Println("----------------------")
	numero := 10
	numero++ //numero = numero + 1
	numero += 15 //numero = numero+ 15
	fmt.Println(numero)

	numero--
	numero-= 20 //numero = numero - 20
	numero *= 3//numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	// FIM DOS PERADORES UNÁRIOS


	

}
