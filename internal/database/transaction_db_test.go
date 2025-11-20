package database

import (
	"database/sql"
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	accountDB     *AccountDB
	clientDB      *ClientDB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
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

	_, err = s.db.Exec(`
	CREATE TABLE transactions (
		id VARCHAR(255) PRIMARY KEY,
		account_id_from VARCHAR(255),
		account_id_to VARCHAR(255),
		amount FLOAT,
		created_at DATE,
		FOREIGN KEY (account_id_from) REFERENCES accounts(id),
		FOREIGN KEY (account_id_to) REFERENCES accounts(id)
	);
	`)
	s.Require().NoError(err)

	s.transactionDB = NewTransactionDB(s.db)
	s.accountDB = NewAccountDB(s.db)
	s.clientDB = NewClientDB(s.db)

	s.client1, _ = entity.NewClient("Charlie", "charlie@example.com")
	s.client2, _ = entity.NewClient("Dana", "dana@example.com")

	accountFrom := entity.NewAccount(s.client1)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	s.db.Close()
	s.db.Exec("DROP TABLE transactions;")
	s.db.Exec("DROP TABLE accounts;")
	s.db.Exec("DROP TABLE clients;")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreateTransaction() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Require().NoError(err)
	err = s.transactionDB.Create(transaction)
	s.Require().NoError(err)
}
