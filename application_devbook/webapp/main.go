package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp_mod/src/router"
	"webapp_mod/src/utils"
)

func main() {

	utils.CarregarTemplates()

	router := router.Gerar()

	fmt.Println("Rodando Webapp na porta 7000...")
	log.Fatal(http.ListenAndServe(":7000", router))

}
