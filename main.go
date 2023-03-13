package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pchchv/golog"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Email    string
	PassHash string
	Balance  float64
}

var testURL string

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		golog.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		golog.Panic("Value %v does not exist", v)
	}
	return value
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createUser(jsonMap map[string]interface{}) (*User, error) {
	var err error
	user := new(User)

	user.Id = uuid.New()
	user.Name = fmt.Sprint(jsonMap["name"])
	user.Email = fmt.Sprint(jsonMap["email"])
	user.PassHash, err = hashPassword(fmt.Sprint(jsonMap["password"]))
	if err != nil {
		return nil, err
	}
	user.Balance = 0

	// TODO: Add a user to the database

	return user, nil
}

func updateBalance(jsonMap map[string]interface{}) (User, error) {
	var user User
	_ = fmt.Sprint(jsonMap["id"])
	_ = fmt.Sprint(jsonMap["operation"]) // Deposit or withdrawal
	_ = fmt.Sprint(jsonMap["amount"])

	// TODO: Retrieve user data from the database by ID
	//       Update user balance
	//       Load the updated user data into the database
	return user, nil
}

func deleteUser(id string) error {
	// TODO: Dlete user from the database

	return nil
}

func getUser(id string) (User, error) {
	var user User

	// TODO: Get user data from the database

	return user, nil
}

func main() {
	server()
}
