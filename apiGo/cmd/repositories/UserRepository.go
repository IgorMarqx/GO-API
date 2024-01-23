package repositories

import (
	"apiGo/cmd/database"
	"apiGo/cmd/model"
)

func SaveUser(user model.User) (bool, error) {
	db := database.DB

	sql, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")

	if err != nil {
		return false, err
	}

	_, err = sql.Exec(user.Name, user.Email)

	if err != nil {
		return false, err
	}

	return true, nil
}
