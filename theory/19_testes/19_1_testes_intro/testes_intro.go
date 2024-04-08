package main

import (
	"fmt"
	enderecos "go_tests_module/19_1_1_enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Avenida 25 de Julho")
	fmt.Println(tipoEndereco)

}
