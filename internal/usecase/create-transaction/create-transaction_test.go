package createtransaction

import (
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john.doe@example.com")
	account1 := entity.NewAccount(client1)
	account1.Deposit(1000)

	client2, _ := entity.NewClient("Jane Smith", "jane.smith@example.com")
	account2 := entity.NewAccount(client2)
	account2.Deposit(500)

	accountGatewayMock := new(AccountGatewayMock)
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	transactionGatewayMock := new(TransactionGatewayMock)
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	useCase := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        200,
	}

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)

	accountGatewayMock.AssertCalled(t, "FindByID", account1.ID)
	accountGatewayMock.AssertCalled(t, "FindByID", account2.ID)
	transactionGatewayMock.AssertCalled(t, "Create", mock.MatchedBy(func(transaction *entity.Transaction) bool {
		return transaction.AccountFrom.ID == account1.ID &&
			transaction.AccountTo.ID == account2.ID &&
			transaction.Amount == 200
	}))
}
