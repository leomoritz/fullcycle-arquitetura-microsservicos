package createtransaction

import (
	"context"
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/event"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/usecase/mocks"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john.doe@example.com")
	account1 := entity.NewAccount(client1)
	account1.Deposit(1000)

	client2, _ := entity.NewClient("Jane Smith", "jane.smith@example.com")
	account2 := entity.NewAccount(client2)
	account2.Deposit(500)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreatedEvent()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	useCase := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        200,
	}

	output, err := useCase.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
