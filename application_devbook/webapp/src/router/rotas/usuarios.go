package rotas

import (
	"net/http"
	"webapp_mod/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Nome:               "Página de Criar Usuário",
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCadastroDeUsuário,
		RequerAutenticacao: false,
	},
}
