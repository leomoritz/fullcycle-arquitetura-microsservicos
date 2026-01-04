package createtransaction

import (
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/gateway"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
	eventDispatcher    events.EventDispatcherInterface
	transactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: transactionGateway,
		accountGateway:     accountGateway,
		eventDispatcher:    eventDispatcher,
		transactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.accountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.accountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	uc.transactionCreated.SetPayload(output)
	uc.eventDispatcher.Dispatch(uc.transactionCreated)

	return output, nil
}
