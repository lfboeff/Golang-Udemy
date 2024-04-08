package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários."))
}

func main() {
	// HTTP é um protocolo de comunicação, base da comunicação web

	// HTTP funciona no modelo cliente <=> servidor

	// Cliente faz requisição ("Request") > servidor processa a requisição e devolve uma resposta ("Response")

	// Rotas
	// - URI: identificador do recurso
	// - Método: GET (ler dados existentes), POST (criar novos dados), PUT (atualizar dados existentes), DELETE (remover dados existentes)

	http.HandleFunc("/home", home) // http://localhost:5000/home

	http.HandleFunc("/usuarios", usuarios) // http://localhost:5000/usuarios

	log.Fatal(http.ListenAndServe(":5000", nil))

}
