package token

import (
	database "go-api/db"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}

	return bytes
}

func CheckPassword(providedPassword string, user *database.User) error {
	log.Println("provided", providedPassword)
	log.Println("user in db", user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
