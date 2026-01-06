package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	_, err := NewClient("", "")
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("Jane Doe", "jane@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe", client.Name)
	assert.Equal(t, "jane@j.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe", "")
	assert.NotNil(t, err)
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Len(t, client.Accounts, 1)
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountWhenAccountDoesNotBelongToClient(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	client2, _ := NewClient("Jane Doe", "jane@j.com")
	account := NewAccount(client2)
	err := client1.AddAccount(account)
	assert.NotNil(t, err)
}
