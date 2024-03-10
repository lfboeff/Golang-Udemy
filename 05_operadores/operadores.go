package main

import "fmt"

func main() {

	// Operadores aritméticos:

	soma := 1 + 1
	subtracao := 10 - 6
	multiplicacao := 5 * 5
	divisao := 5 / 5
	restoDivisao := 25 % 5
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	var n1 int16 = 10
	var n2 int32 = 25
	soma_2 := n1 + int16(n2)
	fmt.Println(soma_2)

	// Operadores atribuição:

	fmt.Println()
	var variavel_1 string = "String 1"
	variavel_2 := "String 2"
	fmt.Println(variavel_1, "-", variavel_2)

	// Operadores relacionais:

	fmt.Println()
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores lógicos:

	fmt.Println()
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro)
	fmt.Println(falso || verdadeiro)
	fmt.Println(!falso && verdadeiro)

	// Operadores unários:

	fmt.Println()
	numero := 10
	numero += 1
	numero++
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero *= 5
	fmt.Println(numero)
	numero %= 5
	fmt.Println(numero)

	// Operador ternário:

	fmt.Println()
	// texto := numero > 5 ? "maior que 5" : "menor que 5" >>> não existe isso em Go
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
