package main

import (
	"errors"
	"fmt"
)

func main() {

	// Nuemros inteiros
	// int8, int16, int32, int64
	var numero int = 100000000000
	numero2 := 100000000000

	fmt.Println(numero)
	fmt.Println(numero2)

	//uint sem sinal ... nao aceita colocar sinal na frente Ex: -150

	var numero3 uint32 = 10000
	fmt.Println(numero3)

	//Alias (apelidos)
	// INT32 = RUNE

	var numero4 rune = 123456
	fmt.Println(numero4)

	//BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	// Numeros reais float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.45
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2 "
	fmt.Println(str2)

	// Mais prosximo de um Char (carctere) 1 somente

	char := 'A'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)
	//. erro error = erros (erro: nome da variável, error: do tipo error, errors: Pacote de errors dentro do Go.)
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)


}
