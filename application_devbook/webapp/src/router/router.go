package router

import (
	"webapp_mod/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {

	router := mux.NewRouter()

	return rotas.Configurar(router)

	// return mux.NewRouter()
}
