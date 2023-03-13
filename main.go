package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

var (
	testURL string
	conn    *pgx.Conn
)

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

	user.Name = fmt.Sprint(jsonMap["name"])
	user.Email = fmt.Sprint(jsonMap["email"])
	user.PassHash, err = hashPassword(fmt.Sprint(jsonMap["password"]))
	if err != nil {
		return nil, err
	}
	user.Balance = 0

	row := conn.QueryRow(context.Background(),
		"INSERT INTO users (name, email, password, balance) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.PassHash, user.Balance)
	err = row.Scan(&user.Id)
	if err != nil {
		return nil, fmt.Errorf("Unable to INSERT: %v\n", err)
	}

	return user, nil
}

func updateBalance(jsonMap map[string]interface{}) (User, error) {
	id := fmt.Sprint(jsonMap["id"])
	operation := fmt.Sprint(jsonMap["operation"]) // deposit or withdrawal
	user, err := getUser(id)
	if err != nil {
		return user, err
	}
	amount, err := strconv.ParseFloat(fmt.Sprint(jsonMap["amount"]), 64)
	if err != nil {
		return user, err
	}

	switch operation {
	case "deposit":
		user.Balance += amount
	case "withdrawal":
		user.Balance -= amount
	}

	_, err = conn.Exec(context.Background(), "update users set balance=$1 where id=$2", user.Balance, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func deleteUser(id uuid.UUID) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Unable to DELETE: %v\n", err)
	}
	return nil
}

func getUser(id string) (User, error) {
	var user User

	uid, err := uuid.Parse(id)
	if err != nil {
		return user, err
	}

	err = conn.QueryRow(context.Background(), "select user from users where id=$1", uid).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func main() {
	db()
	server()
}
