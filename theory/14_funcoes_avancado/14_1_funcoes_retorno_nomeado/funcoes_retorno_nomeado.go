package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subt int) {
	soma = n1 + n2
	subt = n1 - n2
	return
}

func main() {
	resultadoSoma, resultadoSubt := calculosMatematicos(10, 3)
	fmt.Println(resultadoSoma, resultadoSubt)
}
