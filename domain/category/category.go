package category

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)



type UseCase interface{
	Create(m *model.Category) error
	Update(m *model.Category) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Category, error)
	GetAll()(model.Categorys, error)

	UpdateEsceptImage(ID uuid.UUID, updatedCategory model.Category) error
}

type Storage interface {
	Create(m *model.Category) error
	Update(m *model.Category) error
	Delete(ID uuid.UUID) error


	GetByID(ID uuid.UUID) (model.Category, error)
	GetAll() (model.Categorys, error)

	UpdateEsceptImage(ID uuid.UUID, updatedCategory model.Category) error
}
