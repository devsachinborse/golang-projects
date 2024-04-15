package main

import (
	"math/rand"
	"time"
)

type createAccountReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID         int       `json:"id"`
	FistName   string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Number     int64     `json:"number"`
	Balance    int64     `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		//ID:         rand.Intn(1000),
		FistName:   firstName,
		LastName:   lastName,
		Number:     int64(rand.Intn(100000)),
		Created_at: time.Now().UTC(),
	}
}
