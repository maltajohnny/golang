package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Arquivo structs")

	enderecoExemplo := endereco{"Rua dos Bobos", 0}
	var u usuario
	u.nome = "Davi"
	u.idade = 18
	fmt.Println(u)

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)

}
