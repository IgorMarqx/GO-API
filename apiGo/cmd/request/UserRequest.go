package request

import (
	"apiGo/cmd/model"
	"errors"
	"regexp"
)

func Validate(user model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidaEmail(user.Email) {
		return errors.New("invalid format email")
	}

	return nil
}

func isValidaEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
