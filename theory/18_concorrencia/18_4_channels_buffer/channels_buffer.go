package main

import "fmt"

func main() {

	// canais com buffer são utilizados quando queremos utilizar channels dentro de uma mesma função, tanto para enviar quanto para receber um dado

	canal := make(chan string, 2) // canal com capacidade de 2 dados a serem recebidos em canal antes do programa travar

	canal <- "Olá mundo!"
	canal <- "Hey"
	// canal <- "Oops" // isso faria com que o programa desse erro de deadlock, pois não há nada no programa esperando por uma 3a chamada de canal (o buffer está com valor de 2 apenas)

	mensagem1 := <-canal
	fmt.Println(mensagem1)

	mensagem2 := <-canal
	fmt.Println(mensagem2)

}
