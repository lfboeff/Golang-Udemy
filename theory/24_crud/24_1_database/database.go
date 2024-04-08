package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // driver de conexão com mysql
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		fmt.Println("Erro de abertura do banco de dados.")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Erro de ping.")
		return nil, err
	}

	return db, nil

}
