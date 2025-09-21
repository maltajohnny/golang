package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i na posição: ", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J na posição", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [ ]string{"João", "Davi", "Lucas"}
	// for _, nome := range nomes {
	// 	fmt.Println( nome)
	// 	time.Sleep(time.Second)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// 	time.Sleep(time.Second)

	// }

	// usuario := map[string]string{
	// 	"Nome":       " Leonardo",
	// 	"Sobrenome ": " Silva",
	// }
	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }

	for {
		fmt.Println("Loop Infinito")
		time.Sleep(time.Second)
	}
}
