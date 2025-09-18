package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {

	return n1 + n2

}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	var f1 = func() {
		fmt.Println("Texto da função 1")
	}
	f1()

	var f2 = func(str string) {
		fmt.Println(str)
	}
	f2("Texto da função 2")

	var f3 = func(txt string) string {
		// 		fmt.Println(txt) poderia também por junto com o retorno
		return txt
	}
	resultado := f3("Texto da função 3")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// caso queira que exiba somente um resultado eu coloco um underline no lugar de resultadoSubtracao e não passo ele no print.

	

}
