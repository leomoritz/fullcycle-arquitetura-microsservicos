package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/database"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/event"
	createaccount "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-account"
	createclient "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-client"
	createtransaction "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-transaction"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost:3307", "walletcore"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	dispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	// dispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, dispatcher, transactionCreatedEvent)

}
