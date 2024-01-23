package request

import (
	"apiGo/cmd/model"
	"regexp"
)

type ValidationResult struct {
	Error   bool
	Message string
}

func Validate(user model.User) ValidationResult {
	if user.Name == "" {
		return ValidationResult{Error: true, Message: "name is required"}
	}

	if user.Email == "" {
		return ValidationResult{Error: true, Message: "email is required"}
	}

	if !isValidaEmail(user.Email) {
		return ValidationResult{Error: true, Message: "email is invalid"}
	}

	return ValidationResult{Error: false, Message: ""}
}

func isValidaEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
