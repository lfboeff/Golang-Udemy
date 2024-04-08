package server

import (
	db_connect "crud_mod/24_1_db_connect"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	var user usuario

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	db, err := db_connect.Conectar()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer db.Close()

	// INSERT INTO usuarios (name, email) VALUES (user.Name, user.Email)
	statement, err := db.Prepare("insert into usuarios (name, email) values (?, ?)") // PREPARE STATEMENT: evitar um ataque chamado de "SQL injection"
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	// STATUS CREATED = 201
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

	/*	POSTMAN:
		POST
		localhost:5000/usuarios
		{
			"name": "João",
			"email": "joao@gmail.com"
		}
	*/

}

// BuscarUsuarios traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, err := db_connect.Conectar()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer db.Close()

	// SELECT * FROM usuarios
	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer rows.Close()

	var users []usuario

	for rows.Next() {
		var user usuario
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte(error.Error(err)))
			return
		}
		users = append(users, user)
	}

	// STATUS OK = 200
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	/*	POSTMAN:
		GET
		localhost:5000/usuarios
	*/
}

// BuscarUsuario traz um usuário específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	// fmt.Println(ID)

	db, err := db_connect.Conectar()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer db.Close()

	// SELECT * FROM usuarios WHERE id = ?
	row, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer row.Close()

	var user usuario

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte(error.Error(err)))
			return
		}
	}

	// STATUS OK = 200
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	/*	POSTMAN:
		GET
		localhost:5000/usuarios/1
	*/

}

// AtualizarUsuario altera os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	var user usuario

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	db, err := db_connect.Conectar()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer db.Close()

	// UPDATE usuarios SET name = ?, email = ? WHERE id = ?
	statement, err := db.Prepare("update usuarios set name = ?, email = ? where id = ?") // PREPARE STATEMENT: evitar um ataque chamado de "SQL injection"
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	// STATUS OK = 204
	w.WriteHeader(http.StatusNoContent)

	/*	POSTMAN:
		PUT
		localhost:5000/usuarios/3
		{
		    "name": "Daniela",
		    "email": "daniela@gmail.com"
		}
	*/

}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	db, err := db_connect.Conectar()
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer db.Close()

	// DELETE FROM usuarios WHERE id = ?
	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		w.Write([]byte(error.Error(err)))
		return
	}

	// STATUS 204
	w.WriteHeader(http.StatusNoContent)

}
