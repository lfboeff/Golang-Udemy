package controllers

import (
	"api_mod/src/autenticacao"
	"api_mod/src/database"
	"api_mod/src/modelos"
	"api_mod/src/repositorios"
	"api_mod/src/respostas"
	"api_mod/src/seguranca"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Criando um usuário..."))

	bodyRequest, err := io.ReadAll(r.Body)

	// err = errors.New("fake error")
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		// The 422 (Unprocessable Content) status code indicates that the server understands the content type of the
		// request content, and the syntax of the request content is correct, but it was unable to process the contained
		// instructions. For example, this status code can be sent if an XML request content contains well-formed
		// (i.e., syntactically correct), but semantically erroneous XML instructions.
		return
	}

	var usuario modelos.Usuario

	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		// The 400 (Bad Request) status code indicates that the server cannot or will not process the request due
		// to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message
		// framing, or deceptive request routing)
		return
	}

	if err = usuario.Preparar("cadastro"); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		// The 500 (Internal Server Error) status code indicates that the server encountered an unexpected condition
		// that prevented it from fulfilling the request.
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	usuarioId, err := repositorioUsuarios.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuario.ID = usuarioId
	respostas.JSON(w, http.StatusCreated, usuario)
	// The 201 (Created) status code indicates that the request has been fulfilled and has resulted in one or more new
	// resources being created. The primary resource created by the request is identified by either a Location header
	// field in the response or, if no Location header field is received, by the target URI

	// w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioId)))

	/* POSTMAN:
		POST
		localhost:5000/usuarios
		{
	    "nome": "Felipe",
	    "nick": "lfb",
	    "email": "felipe@gmail.com",
	    "senha": "1234567"
		}
	*/
}

// BuscarUsuarios busca todos os usuários do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Buscando todos os usuários..."))

	var nomeOuNick = strings.ToLower(r.URL.Query().Get("usuario"))
	// fmt.Println(nomeOuNick)

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	usuarios, err := repositorioUsuarios.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

	// localhost:5000/usuarios?usuario=luis

}

// BuscarUsuario busca um usuário do banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Buscando um usuário..."))

	parameters := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
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

	usuario, err := repositorioUsuarios.BuscarPorID(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// AtualizarUsuario altera as informações de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	//	w.Write([]byte("Atualizando um usuário..."))

	parameters := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

	fmt.Println(usuarioIdNoToken)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario

	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
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

	if err = repositorioUsuarios.Atualizar(usuarioID, usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Deletando um usuário..."))
	parameters := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	if err = repositorioUsuarios.Deletar(usuarioID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// SeguirUsuario permite que um usuário siga um outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {

	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	usuarioASeguirID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioASeguirID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	if err = repositorioUsuarios.Seguir(usuarioASeguirID, seguidorID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite que um usuário pare de seguir um outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {

	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	usuarioPararSeguirID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioPararSeguirID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioUsuarios := repositorios.NovoRepositorioDeUsuarios(db)

	if err = repositorioUsuarios.PararDeSeguir(usuarioPararSeguirID, seguidorID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
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

	seguidores, err := repositorioUsuarios.BuscarSeguidores(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
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

	usuariosSeguidos, err := repositorioUsuarios.BuscarSeguindo(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuariosSeguidos)
}

// AtualizarSenha permite atualizar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	usuarioIDNaRequest, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioIDNoToken != usuarioIDNaRequest {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var senha modelos.Senha

	if err = json.Unmarshal(bodyRequest, &senha); err != nil {
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

	senhaDBComHash, err := repositorioUsuarios.BuscarSenha(usuarioIDNaRequest)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(senhaDBComHash, senha.Atual); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a que está salva no banco"))
		return
	}

	senhaNovaComHash, err := seguranca.Hash(senha.Nova)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorioUsuarios.AtualizarSenha(usuarioIDNaRequest, string(senhaNovaComHash)); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
