package passgen

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a method for hashing password using bcrypt
func HashPassword(password []byte) (pwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return pwd, err
	}
	return string(hash), nil
}

// ComparePasswords is a method for comparing hashed password and plain text password
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
