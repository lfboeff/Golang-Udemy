package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo")
	}()

	fmt.Println("---")

	func(texto string) {
		fmt.Println(texto)
	}("Hellow world")

	fmt.Println("---")

	mensagem := func(texto string) string {
		return fmt.Sprintf("%s, aloha!", texto)
	}("Hellow world")
	fmt.Println(mensagem)

}
