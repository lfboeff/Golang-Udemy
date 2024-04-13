package controllers

import (
	"net/http"
	"webapp_mod/src/utils"
)

// CarregarTelaDeLogin vai carregar a página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "login.html", nil)

	// w.Write([]byte("Tela de Login"))

}

// CarregarTelaDeCadastroDeUsuário vai carregar a página de cadastro de usuário
func CarregarTelaDeCadastroDeUsuário(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
