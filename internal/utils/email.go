package utils

import (
	"employee-app/logger"
	"net/mail"
)

func ValidMailAddress(email string) bool {
	validEmail, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	logger.Info(validEmail.Address)
	return true
}
