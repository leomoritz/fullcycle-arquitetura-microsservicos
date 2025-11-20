package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@gmail.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateAccount_NilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestDeposit(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@gmail.com")
	account := NewAccount(client)
	account.Deposit(500)
	assert.Equal(t, 500.0, account.GetBalance())

	account.Deposit(-100)
	assert.Equal(t, 500.0, account.GetBalance())
}

func TestWithdraw(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@gmail.com")
	account := NewAccount(client)
	account.Deposit(500)

	success := account.Withdraw(200)
	assert.True(t, success)
	assert.Equal(t, 300.0, account.GetBalance())

	success = account.Withdraw(400)
	assert.False(t, success)
	assert.Equal(t, 300.0, account.GetBalance())

	success = account.Withdraw(-50)
	assert.False(t, success)
	assert.Equal(t, 300.0, account.GetBalance())
}

func TestGetBalance(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@gmail.com")
	account := NewAccount(client)
	account.Deposit(500)
	assert.Equal(t, 500.0, account.GetBalance())
}
