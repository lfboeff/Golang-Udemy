package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

func main() {
	soma := somar(2, 5)
	fmt.Println("soma = ", soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt + " aloha"
	}
	resultado := f("Texto da função f(txt)")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(10, 5)
	fmt.Println(resultadoSoma, "-", resultadoSubtracao, "-", resultadoMultiplicacao, "-", resultadoDivisao)

	resultadoSoma, _, _, _ = calculosMatematicos(100, 2)
	fmt.Println(resultadoSoma)
}
