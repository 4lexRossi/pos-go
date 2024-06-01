package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	// initJson := []byte(`{"number":2, "balance":200}`)
	initJson := []byte(`{"n":2, "b":200}`)
	var accountX Account
	err = json.Unmarshal(initJson, &accountX)
	if err != nil {
		panic(err)
	}
	println(accountX.Balance)
}
