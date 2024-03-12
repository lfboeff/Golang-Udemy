package main

import "fmt"

func closure() func() {
	mensagem := "Dentro da função closure()"

	funcao := func() {
		fmt.Println(mensagem)
	}

	return funcao
}

func main() {

	texto := "Dentro da função main()"
	fmt.Println(texto)

	closure()

	fmt.Println("---")

	funcaoNova := closure()
	funcaoNova()

}
