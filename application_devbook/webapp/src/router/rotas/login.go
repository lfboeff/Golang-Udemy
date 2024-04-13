package rotas

import (
	"net/http"
	"webapp_mod/src/controllers"
)

var rotasLogin = []Rota{
	{
		Nome:               "Login Page 1",
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		Nome:               "Login Page 2",
		URI:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
}
