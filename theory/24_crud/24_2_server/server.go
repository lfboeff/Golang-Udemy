package server

import (
	database "crud_mod/24_1_database"
	"encoding/json"
	"fmt"
	"io" // ioutil
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"` // por padrão, tags em json são escritas em minúsculo
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário novo no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter dados de usuário para struct."))
		return
	}

	fmt.Println(usuario)

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	// INSERT INTO usuarios (name, email)) values("name", "email")

	statement, err := db.Prepare("insert into usuarios (name, email) values (?, ?)") // estratégia para evitar ataque "sql injection"
	if err != nil {
		w.Write([]byte("Erro ao criar statement de insert."))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Name, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement de insert."))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o ID do usuário inserido."))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated) // 201 = usuário criado com sucesso!

	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %v", idInserido)))

}
