package database

import (
	"database/sql"
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	var err error
	s.db, err = sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)

	s.clientDB = NewClientDB(s.db)

	_, err = s.db.Exec(`
	CREATE TABLE clients (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255),
		created_at DATE
	);
	`)
	s.Require().NoError(err)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	s.db.Close()
	s.db.Exec("DROP TABLE clients;")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSaveAndGetClient() {
	client, err := entity.NewClient("Alice", "alice@example.com")
	s.Require().NoError(err)

	err = s.clientDB.Save(client)
	s.Require().NoError(err)

	retrievedClient, err := s.clientDB.Get(client.ID)
	s.Require().NoError(err)
	s.Equal(client.ID, retrievedClient.ID)
	s.Equal(client.Name, retrievedClient.Name)
	s.Equal(client.Email, retrievedClient.Email)
	s.WithinDuration(client.CreatedAt, retrievedClient.CreatedAt, 1e9)
}
