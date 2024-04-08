package db_connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // "import implícito"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		fmt.Println("Erro de sql.Open()")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Erro de db.Ping()")
		return nil, err
	}

	return db, nil

}
