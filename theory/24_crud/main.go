package main

import (
	server "crud_mod/24_2_server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// CRUD: CREATE, READ, UPDATE, DELETE
	// CREATE = POST
	// READ = GET
	// UPDATE = PUT
	// DELETE = DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost) // poderia ser "POST" aqui simplesmente

	fmt.Println("Escutando a porta 5000.")
	log.Fatal(http.ListenAndServe(":5000", router))

}

/*
{
    "name": "João",
    "email": "joao.batista@gmail.com"
}
*/
