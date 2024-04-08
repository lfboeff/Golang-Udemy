package main

import "fmt"

func main() {

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")

	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número está entre -10 e 0")
	}
}
