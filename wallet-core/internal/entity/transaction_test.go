package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactional(t *testing.T) {
	client1 := &Client{ID: "client-123", Name: "John Doe"}
	account1 := NewAccount(client1)
	client2 := &Client{ID: "client-456", Name: "Jane Smith"}
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.Equal(t, account1.ID, transaction.AccountFrom.ID)
	assert.Equal(t, account2.ID, transaction.AccountTo.ID)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestExecuteTransaction(t *testing.T) {
	client1 := &Client{ID: "client-123", Name: "John Doe"}
	account1 := NewAccount(client1)
	client2 := &Client{ID: "client-456", Name: "Jane Smith"}
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(500)

	transaction, err := NewTransaction(account1, account2, 200)
	assert.Nil(t, err)

	err = transaction.Execute()
	assert.Nil(t, err)
	assert.Equal(t, 800.0, account1.Balance)
	assert.Equal(t, 700.0, account2.Balance)
}

func TestExecuteTransaction_InsufficientFunds(t *testing.T) {
	client1 := &Client{ID: "client-123", Name: "John Doe"}
	account1 := NewAccount(client1)
	client2 := &Client{ID: "client-456", Name: "Jane Smith"}
	account2 := NewAccount(client2)

	account1.Deposit(100)
	account2.Deposit(500)

	transaction, err := NewTransaction(account1, account2, 200)
	assert.Nil(t, err)

	err = transaction.Execute()
	assert.NotNil(t, err)
	assert.Equal(t, "insufficient funds in the source account", err.Error())
	assert.Equal(t, 100.0, account1.Balance)
	assert.Equal(t, 500.0, account2.Balance)
}

func TestCreateTransaction_InvalidAmount(t *testing.T) {
	client1 := &Client{ID: "client-123", Name: "John Doe"}
	account1 := NewAccount(client1)
	client2 := &Client{ID: "client-456", Name: "Jane Smith"}
	account2 := NewAccount(client2)

	_, err := NewTransaction(account1, account2, -50)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}

func TestCreateTransaction_NilAccounts(t *testing.T) {
	client1 := &Client{ID: "client-123", Name: "John Doe"}
	account1 := NewAccount(client1)

	_, err := NewTransaction(account1, nil, 100)
	assert.NotNil(t, err)
	assert.Equal(t, "both accounts must be provided", err.Error())

	_, err = NewTransaction(nil, account1, 100)
	assert.NotNil(t, err)
	assert.Equal(t, "both accounts must be provided", err.Error())
}
