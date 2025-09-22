package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", user.nome)

}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18

}

func (user *usuario) fazerAniversario(){
	user.idade++

}
func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Usuário2", 30}
	usuario2.salvar()
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)
	

	

}
