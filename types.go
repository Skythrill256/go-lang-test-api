package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
		balance:   0,
	}
}

type createAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
