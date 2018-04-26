package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Rank     int    `json:"rank"`
}

func main() {
	account := &Account{
		Email:    "sample@example.com",
		Password: "password",
		Rank:     1,
	}

	// encode
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(account)
	fmt.Println(buf.String())

	// decode
	decoded := new(Account)
	json.NewDecoder(buf).Decode(decoded)
	fmt.Println(decoded)
}
