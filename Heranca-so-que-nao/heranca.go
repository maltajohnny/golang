package main

import "fmt"

type pessoa struct {
	nome, sobrenome string
	idade           int8
	altura          int16
}

type estudante struct {
	pessoa
	curso, faculdade string
}

func main() {

	fmt.Println("Heranca")
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia de Software", "Unicesumar"}

	fmt.Println(e1)

}
