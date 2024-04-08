package main

import (
	server "crud_mod/24_2_server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD: CREATE, READ, UPDATE, DELETE
// CREATE => POST
// READ => GET
// UPDATE => PUT
// DELETE => DELETE

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)          // .Methods("POST")
	router.HandleFunc("/usuarios", server.BuscarUsuarios).Methods(http.MethodGet)         // .Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.BuscarUsuario).Methods(http.MethodGet)     // .Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuario).Methods(http.MethodPut)  // .Methods("PUT")
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete) // .Methods("DELETE")

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
