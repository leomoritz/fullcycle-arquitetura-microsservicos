package handler

import (
	"fmt"
	"sync"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafkaProducer *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafkaProducer,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, waitGroup *sync.WaitGroup) error {
	defer waitGroup.Done()
	err := h.Kafka.Publish(message, nil, "transactions") // mensagem, chave e nome do tópico
	if err != nil {
		return err
	}

	fmt.Println("Transaction created event published to Kafka topic 'transactions' with data:", message.GetPayload())
	return nil
}
