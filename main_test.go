package main

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pass := "qewhy#fcu3!rt"
	_, err := hashPassword(pass)
	if err != nil {
		t.Fatal()
	}
}

func TestCheckPasswordHash(t *testing.T) {
	pass := "qewhy#fcu3!rt"
	hash, err := hashPassword(pass)
	if err != nil {
		t.Fatal()
	}
	if !checkPasswordHash(pass, hash) {
		t.Fatal()
	}
}

func TestCreateUser(t *testing.T) {
	jsonMap := make(map[string]interface{})
	email := "ipchchv@gmail.com"
	name := "Jack"
	password := "3223414r"

	jsonMap["email"] = email
	jsonMap["name"] = name
	jsonMap["password"] = password

	user, err := createUser(jsonMap)
	if err != nil {
		t.Fatal()
	}

	if user.Email != email {
		t.Fatal()
	}
	if user.Balance != 0 {
		t.Fatal()
	}
	if user.Name != name {
		t.Fatal()
	}
	if !checkPasswordHash(password, user.PassHash) {
		t.Fatal()
	}
}
