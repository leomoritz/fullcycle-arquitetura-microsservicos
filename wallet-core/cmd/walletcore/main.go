package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/database"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/event"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/event/handler"
	createaccount "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-account"
	createclient "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-client"
	createtransaction "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/create-transaction"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/web"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/web/webserver"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/kafka"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/uow"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	dispatcher := events.NewEventDispatcher()
	dispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	dispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, dispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	acountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", acountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
