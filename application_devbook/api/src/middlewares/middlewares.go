package middlewares

import (
	"api_mod/src/autenticacao"
	"api_mod/src/respostas"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println("Autenticando...")

		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}

		proximaFuncao(w, r)
	}
}
