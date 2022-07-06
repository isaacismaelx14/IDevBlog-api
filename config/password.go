package config

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePass(pass string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hash
}

func ValidatePass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}
	return true
}
