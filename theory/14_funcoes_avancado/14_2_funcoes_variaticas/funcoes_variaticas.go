package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

// só podemos ter um parâmetro variático por função, e este deve ser sempre o último

func main() {

	fmt.Println(soma(0, 1, 2, 3, 4))

	fmt.Println("---")

	escrever("Olá mundo!", 7, 6, 5)

}
