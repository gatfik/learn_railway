package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword ...
func HashPassword(p string) (res string, err error) {
	salt := 9
	pass := []byte(p)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, err := bcrypt.GenerateFromPassword(pass, salt)
	if err != nil {
		log.Println(err.Error())
		return
	}

	res = string(passwordHash)
	return
}

// ComparePassword ...
func ComparePassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}
	return true
}
