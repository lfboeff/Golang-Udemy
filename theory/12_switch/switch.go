package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		fmt.Println("Oops!")
		return "Número inválido"
	}
}

func diaDaSemana3(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// fallthrough >>> isso faria com que, se numero == 1, a condição de numero == 2 seria executada também
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		fmt.Println("Oops!")
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {

	dia := diaDaSemana(7)
	fmt.Println(dia)

	fmt.Println()
	dia2 := diaDaSemana2(15)
	fmt.Println(dia2)

	fmt.Println()
	dia3 := diaDaSemana3(3)
	fmt.Println(dia3)

}
