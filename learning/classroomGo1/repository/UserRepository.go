package repository

import (
	"classroomGo1/db"
	"database/sql"
	"errors"
)

type Queryable interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func SaveUser(name, email string) (bool, error) {
	database := db.DB

	// DONE - Validar se o usuario e o email está preenchido
	validateUser := UserRequest(name, email)

	if !validateUser {
		return false, errors.New("O usuário e o email são obrigatórios")
	}

	// DONE - Verificar se o email já existe no banco de dados
	verifyEmail, err := verifyIfEmailExist(name, email, database)

	if !verifyEmail {
		return false, err
	}

	// DONE - Inserir o usuário no banco de dados
	query, err := database.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")

	if err != nil {
		return false, err
	}
	defer query.Close()

	_, err = query.Exec(name, email)
	if err != nil {
		return false, err
	}

	// DONE - Retornar true ou false
	return true, nil
}

func UserRequest(name, email string) bool {
	if name == "" || email == "" {
		return false
	}

	return true
}

func verifyIfEmailExist(name, email string, database Queryable) (bool, error) {
	rows, err := database.Query("SELECT id FROM users WHERE email = ?", email)

	if err != nil {
		return false, errors.New("Problema com a conexão com o banco de dados")
	}
	defer rows.Close()

	if rows.Next() {
		return false, errors.New("O email já existe")
	}

	return true, nil
}
