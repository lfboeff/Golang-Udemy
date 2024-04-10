package controllers

import (
	"api_mod/src/autenticacao"
	"api_mod/src/database"
	"api_mod/src/modelos"
	"api_mod/src/repositorios"
	"api_mod/src/respostas"
	"api_mod/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario

	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	usuarioSalvoNoBanco, err := repositorioUsuarios.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	// w.Write([]byte("Você está logado, parabéns!"))

	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))

}
