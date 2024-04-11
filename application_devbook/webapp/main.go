package main

import (
	"fmt"
	"net/http"
	"webapp_mod/src/router"
)

func main() {

	fmt.Println("Rodando Webapp na porta 3000...")

	router := router.Gerar()

	http.ListenAndServe(":3000", router)

}
