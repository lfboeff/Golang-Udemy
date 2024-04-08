package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Início do programa")

	canal := escrever("Olá mundo!") // "canal" é do tipo "<-chan", ou seja, só recebe dados

	fmt.Println("Passagem por canal")

	for i := 0; i < 10; i++ {
		fmt.Println("Entrada no for de main()")
		fmt.Println(<-canal)
		fmt.Println("Saída do for de main()")
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string) <-chan string {

	fmt.Println("Entrada em escrever()")

	canal := make(chan string)

	go func() {
		i := 0
		for {

			fmt.Println("Entrada no for de escrever()")

			canal <- fmt.Sprintf("Valor recebido: %s - %v", texto, i)
			time.Sleep(time.Second)
			i++

			fmt.Println("Saída do for de escrever()")

		}
	}()

	fmt.Println("Saída de escrever()")

	return canal

}
