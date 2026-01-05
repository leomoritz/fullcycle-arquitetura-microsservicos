package createaccount

import (
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"
	"github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	clientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: accountGateway,
		clientGateway:  clientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.clientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.accountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	output := &CreateAccountOutputDTO{
		ID: account.ID,
	}

	return output, nil
}
