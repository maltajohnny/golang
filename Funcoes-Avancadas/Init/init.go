package main

import "fmt"

var n int

func init() {
	fmt.Println("Função Init sendo Executada")
	n = 10
}

func main() {
	fmt.Println(n)
	fmt.Println("Função main sendo executada")
}
