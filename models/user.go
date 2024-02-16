package models

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type UserCred struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Encrypt encrypts user passowrds for storage
func (user *UserCred) Encrypt(password string) {
	hash, bcryptErr := bcrypt.GenerateFromPassword([]byte(password), 12)
	if bcryptErr != nil {
		log.Fatal(bcryptErr)
	}

	user.Password = string(hash)
}

// GetSignedJWT returns a signed token containg the users _id and the expiration time
func (user *UserCred) GetSignedJWT(id string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte("RISHI"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
