package category

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type Category struct {
	storage Storage
}

func New(s Storage) Category {
	return Category{storage: s}
}

func (c Category) Create(m *model.Category) error {
	ID, err := uuid.NewUUID()
	if err != nil { return fmt.Errorf("%s %w", "uuid.NewUUID()", err) }
	m.ID = ID

	if len(m.Images) == 0 { m.Images = ""}

	m.CreatedAt = time.Now().Unix()

	err = c.storage.Create(m)
	if err != nil { return err }
	return nil
}

func (c Category) Update(m *model.Category)error {
	if !m.HasID() { return fmt.Errorf("Update HasID") }

	//if len(m.Images) == 0 { m.Images = []byte(`{}`) }
	if len(m.Images) == 0 { m.Images = ""}

	m.UpdatedAt = time.Now().Unix()

	err := c.storage.Update(m)
	if err != nil { return err }

	return nil
}

func (c Category) UpdateEsceptImage(ID uuid.UUID, updatedCategory model.Category) error {
    // Obtener la categoría por su ID
    m, err := c.GetByID(ID)
    if err != nil {
        return err
    }

    // Actualizar todos los campos excepto la imagen
    m.Nombre = updatedCategory.Nombre
    m.Description = updatedCategory.Description
    m.Activo = updatedCategory.Activo
    // Actualiza otros campos según la estructura de tu modelo

    m.UpdatedAt = time.Now().Unix()

    // Actualizar la categoría en el almacenamiento
    err = c.storage.Update(&m) // Aquí se corrige el paso del puntero
    if err != nil {
        return err
    }
    return nil
}


func (c Category) Delete(ID uuid.UUID) error {
	err := c.storage.Delete(ID)
	if err != nil { return err}

	return nil
}

func (c Category) GetByID(ID uuid.UUID) (model.Category, error) {
	category, err := c.storage.GetByID(ID)
	if err != nil { return model.Category{}, fmt.Errorf("category: %w", err) }

	return category, nil
}

func (c Category) GetAll() (model.Categorys, error) {
	categorys, err := c.storage.GetAll()
	if err != nil { return model.Categorys{}, fmt.Errorf("Category: %w", err) }

	return categorys, nil
}
