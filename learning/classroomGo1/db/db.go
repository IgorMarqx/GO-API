package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// INFO -  DB é uma variável global que armazenará a conexão com o banco de dados
var DB *sql.DB

func InitDB() {
	connection := "root:root@tcp(127.0.0.1:3306)/golearning"

	// INFO -  Abre uma nova conexão com o banco de dados
	db, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}

	// INFO - Testa a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Conexão com o banco de dados aberta com sucesso!")

	DB = db
}
