package gateway

import "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
