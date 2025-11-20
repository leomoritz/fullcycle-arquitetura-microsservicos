package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom == nil || t.AccountTo == nil {
		return errors.New("both accounts must be provided")
	}
	return nil
}

func (t *Transaction) Execute() error {
	var amount = t.Amount

	if t.AccountFrom.Balance < amount {
		return errors.New("insufficient funds in the source account")
	}

	t.AccountFrom.Withdraw(amount)
	t.AccountTo.Deposit(amount)
	return nil
}
