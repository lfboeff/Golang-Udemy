package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// fmt.Println(stringConexao)

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		fmt.Println("Erro de Open...")
		log.Fatal(err)
	}
	defer db.Close()
	// fmt.Println(db)

	if err = db.Ping(); err != nil {
		fmt.Println("Erro de Ping...")
		log.Fatal(err)
	}
	fmt.Println("Conexão está aberta!")

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		fmt.Println("Erro de Query")
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows)

}
