package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(rawPass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePassword(hashedPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
