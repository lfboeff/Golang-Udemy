package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	fmt.Println(variavel1)

	variavel2 := "variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "aloha"
		variavel4 string = "irmãos"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "eaí", "galera"
	fmt.Println(variavel5, variavel6)

	const contante1 string = "valor estático"
	fmt.Println(contante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
