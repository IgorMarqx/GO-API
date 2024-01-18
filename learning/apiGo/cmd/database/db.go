package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	connection := "root:root@tcp(127.0.0.1:3306)/golearning"
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar ao banco de dados: %s", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	DB = db
}
