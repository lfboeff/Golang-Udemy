package rotas

import (
	"api_mod/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Nome:               "Login",
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
