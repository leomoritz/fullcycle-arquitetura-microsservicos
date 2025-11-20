package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}

	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		return
	}
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Withdraw(amount float64) bool {
	if amount <= 0 || amount > a.Balance {
		return false
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return true
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}
