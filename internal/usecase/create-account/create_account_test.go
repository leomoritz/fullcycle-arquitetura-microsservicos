package createaccount

import (
	"testing"

	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	clientGatewayMock := &ClientGatewayMock{}
	accountGatewayMock := &AccountGatewayMock{}

	client, _ := entity.NewClient("John Doe", "john.doe@example.com")
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)

	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)

	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
