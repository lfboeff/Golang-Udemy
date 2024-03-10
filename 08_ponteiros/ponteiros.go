package main

import "fmt"

func main() {

	var v1 int = 10
	var v2 int = v1
	fmt.Println("v1:", v1, "- v2:", v2)

	v1++
	fmt.Println("v1:", v1, "- v2:", v2)

	// ponteiro é uma referência de memória:

	fmt.Println()
	var v3 int
	var ponteiro *int

	v3 = 50
	ponteiro = &v3
	fmt.Println("v3:", v3, "- ponteiro:", ponteiro, "- valor do ponteiro:", *ponteiro)

	v3 = 75
	fmt.Println("v3:", v3, "- ponteiro:", ponteiro, "- valor do ponteiro:", *ponteiro) // notamos que o endereço de memória não mudou, mas seu valor armazenado sim

}
