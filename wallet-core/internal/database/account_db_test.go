package database

import (
	"database/sql"
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	var err error
	s.db, err = sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)

	_, err = s.db.Exec(`
	CREATE TABLE clients (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255),
		created_at DATE
	);
	`)
	s.Require().NoError(err)

	_, err = s.db.Exec(`
	CREATE TABLE accounts (
		id VARCHAR(255) PRIMARY KEY,
		client_id VARCHAR(255),
		balance FLOAT,
		created_at DATE,
		FOREIGN KEY (client_id) REFERENCES clients(id)
	);
	`)
	s.Require().NoError(err)

	s.accountDB = NewAccountDB(s.db)
	s.client, _ = entity.NewClient("Bob", "bob@example.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	s.db.Close()
	s.db.Exec("DROP TABLE accounts;")
	s.db.Exec("DROP TABLE clients;")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Require().NoError(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	// First, save the client
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	// Now, create and save the account
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Require().NoError(err)

	// Retrieve the account by ID
	retrievedAccount, err := s.accountDB.FindByID(account.ID)
	s.Require().NoError(err)
	s.Equal(account.ID, retrievedAccount.ID)
	s.Equal(account.Client.ID, retrievedAccount.Client.ID)
	s.Equal(account.Balance, retrievedAccount.Balance)
	s.Equal(account.Client.Name, retrievedAccount.Client.Name)
	s.Equal(account.Client.Email, retrievedAccount.Client.Email)
	s.WithinDuration(account.CreatedAt, retrievedAccount.CreatedAt, 1e9)
}
