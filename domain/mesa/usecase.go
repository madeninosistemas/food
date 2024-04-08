package mesa

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)


type Mesa struct{
	storage Storage
}

func New(s Storage) Mesa {
	return Mesa{storage: s}
}

func (me Mesa) Create(m * model.Mesa) error {
	ID, err := uuid.NewUUID()
	if err != nil { return fmt.Errorf("%s %w", "uuid.UUID()", err) }
	m.ID = ID

	if len(m.Images) == 0 { m.Images = "" }

	m.CreatedAt = time.Now().Unix()

	err = me.storage.Create(m)
	if err != nil { return err }

	return nil
}

func (me Mesa) Update(m * model.Mesa) error {
	if !m.HasID(){ return fmt.Errorf("Update HasID") }

	if len(m.Images) == 0 { m.Images = "" }
	m.UpdatedAt = time.Now().Unix()

	err := me.storage.Update(m)
	if err != nil { return err }
	return nil
}

func (me Mesa) Delete(ID uuid.UUID) error {
	err := me.storage.Delete(ID)
	if err != nil { return err }
	return nil
}

func (me Mesa) GetByID(ID uuid.UUID) (model.Mesa, error) {
	mesa, err := me.storage.GetByID(ID)
	if err != nil {
		return model.Mesa{},
		fmt.Errorf("menu: %w", err)
	}
	return mesa, nil
}


func (me Mesa) GetAll()(model.Mesas, error) {
	mesas, err := me.storage.GetAll()
	if err != nil { return model.Mesas{}, fmt.Errorf("Mesa: %w", err) }

	if len(mesas) == 0 {
		return model.Mesas{}, nil
	}

	return mesas, nil
}

