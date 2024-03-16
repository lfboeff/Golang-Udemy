package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		i := 1
		for {
			time.Sleep(time.Second * 1)
			canal1 <- fmt.Sprintf("Canal 1 - %v", i)
			i++
		}
	}()

	go func() {
		i := 1
		for {
			time.Sleep(time.Second * 4)
			canal2 <- fmt.Sprintf("Canal 2 - %v", i)
			i++
		}
	}()

	/*
		for {
			mensagem1 := <-canal1
			fmt.Println(mensagem1)

			mensagem2 := <-canal2
			fmt.Println(mensagem2)
		}
	*/
	for {
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}
}
