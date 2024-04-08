package mesa

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type UseCase interface{
	Create(m *model.Mesa) error
	Update(m *model.Mesa) error
	Delete(ID uuid.UUID) error

	GetByID (ID uuid.UUID) (model.Mesa, error)
	GetAll () (model.Mesas, error)
}

type Storage interface {
	Create(m *model.Mesa) error
	Update(m *model.Mesa) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Mesa, error)
	GetAll() (model.Mesas, error)
}
