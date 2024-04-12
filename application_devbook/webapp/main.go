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

	fmt.Println("Rodando Webapp na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
