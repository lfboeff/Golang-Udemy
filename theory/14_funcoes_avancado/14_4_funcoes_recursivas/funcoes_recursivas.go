package main

import "fmt"

func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {

	// Série de Fibonacci: 1, 1, 2, 3, 5, 8, 13, ...
	posicao := uint(3)
	fmt.Println(fibonacci(posicao))

	fmt.Println("---")

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
