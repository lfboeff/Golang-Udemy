package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	buffer := 45
	tarefas := make(chan int, buffer)
	resultados := make(chan int, buffer)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	duration := time.Since(start)
	fmt.Println(duration) // 10.914596209s // 7.041747249s
}

func worker(tarefas <-chan int, resultados chan<- int) {
	// o argumento "tarefas <-chan int" significa que essa
	// variável, apesar de ser do tipo channel, só irá receber dados
	// já "resultados chan<- int" significa que resultados só
	// receberá dados, apesar de também ser do tipo channel

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int) int {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}
