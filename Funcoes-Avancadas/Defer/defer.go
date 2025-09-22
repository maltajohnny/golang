package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("O aluno está aprovado: ", media)
		return true
	}
	fmt.Println("O aluno está reprovado: ", media)

	return false
}

func main() {
	// fmt.Println("Defer")

	// defer funcao1()
	// DEFER (Fâr) em INglês = Adiar
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
