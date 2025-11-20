package gateway

import "github.com/leomoritz/fullcycle-arquitetura-microsservicos/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
