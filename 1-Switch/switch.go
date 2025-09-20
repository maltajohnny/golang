package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	diaSemana := ""
	switch {
	case numero == 1:
		diaSemana = "Domingo"
	case numero == 2:
		diaSemana = "Segunda"
	case numero ==3:
		diaSemana = "Terça"
	case numero == 4:
		diaSemana = "Quarta"
	case numero == 5:
		diaSemana = "Quinta"
	case numero == 6:
		diaSemana = "Sexta"
	case numero == 7:
		diaSemana = "Sábado"

	default:
		diaSemana = "Número inválido"
	}
	return diaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)
	fmt.Println(dia)
	fmt.Println("--------------------------")


	dia3 := diaDaSemana2(3)
	fmt.Println(dia3)
}
