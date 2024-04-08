package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a função init()")
	n = 10
	fmt.Println("n:", n)
}

// a função init() é sempre executada primeiro (na hora em que o correspondente arquivo for executado)
// ela é bastante utilizada para realizar setups/inicializações

func main() {
	fmt.Println("Função main() sendo executada")
}
