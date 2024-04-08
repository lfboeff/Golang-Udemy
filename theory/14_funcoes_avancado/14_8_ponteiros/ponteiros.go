package main

import "fmt"

func inverterSinal(n int) int {
	return n * (-1)
}

func inverterSinalComPonteiro(n *int) {
	*n *= (-1)
}

func main() {

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numero, "-->", numeroInvertido)

	fmt.Println("---")

	inverterSinalComPonteiro(&numero)
	fmt.Println(numero)

	fmt.Println("---")

	novoNumero := -17
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
