package promocion

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type Storage interface {
	Create(m * model.Promocion) error
	Update(m * model.Promocion) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Promocion, error)
	GetAll() (model.Promociones, error)
}

type UseCase interface {
	Create(m * model.Promocion) error
	Update(m * model.Promocion) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Promocion, error)
	GetAll() (model.Promociones, error)
}


