package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // "import implícito"
)

func main() {
	// "usuario:senha@/banco"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(stringConexao)

	db, err := sql.Open("mysql", stringConexao) // abre a conexão com o banco de dados
	if err != nil {
		fmt.Println("Erro de sql.Open()")
		log.Fatal(err)
	}
	defer db.Close() // garante que a conexão com o banco de dados será fechada quando pararmos de usá-lo

	if err = db.Ping(); err != nil {
		fmt.Println("Erro de db.Ping()")
		log.Fatal(err)
	}
	fmt.Println("A conexão com o banco de dados mysql está aberta!")

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // garante que a leitura de rows será fechada quando pararmos de usá-la

	fmt.Println(rows)

}
