package handler

import (
	"fmt"
	"sync"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/kafka"
)

type UpdateBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewUpdateBalanceKafkaHandler(kafka *kafka.Producer) *UpdateBalanceKafkaHandler {
	return &UpdateBalanceKafkaHandler{
		Kafka: kafka,
	}
}
func (h *UpdateBalanceKafkaHandler) Handle(message events.EventInterface, waitGroup *sync.WaitGroup) error {
	defer waitGroup.Done()
	err := h.Kafka.Publish(message, nil, "balances") // mensagem, chave e nome do tópico
	if err != nil {
		return err
	}

	fmt.Println("Balance updated event published to Kafka topic 'balances' with data:", message.GetPayload())
	return nil
}
