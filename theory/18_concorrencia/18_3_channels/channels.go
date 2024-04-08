package main

import (
	"fmt"
	"time"
)

func main() {

	// channels podem tanto enviar quanto receber dados e são usados para sincronizar goroutines

	canal := make(chan string) // "chan string" significa que esse canal só poderá trafegar dados do tipo string

	go escrever("Olá mundo!", canal)

	fmt.Println("Função escrever() começou a ser executada.")

	/*
		for {
			mensagem, canal_aberto := <-canal // isso significa que estamos recebendo um valor em canal
			if !canal_aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/
	for mensagem := range canal { // esse trecho faz a mesma coisa que o de cima
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	//	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // isso significa que o nosso canal estará recebendo o dado de tipo string "texto"
		time.Sleep(time.Second)
	}
	close(canal) // fechamento do canal para o programa não usá-lo mais, agora que ele não é mais preciso
}

// o erro de deadlock só é notado em execução (go não consegue pegar isso durante a compilação)
