package producto

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type Producto struct {
	storage Storage
}

func New(s Storage) Producto {
	return Producto{storage: s}
}

func (p Producto) Create(m *model.Producto) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID

	if len(m.Imagen) == 0 { m.Imagen = "" }

	m.CreateAt = time.Now().Unix()

	err = p.storage.Create(m)
	if err != nil { return err }

	return nil
}

func (p Producto) Update(m * model.Producto) error {
	if !m.HasID() { return fmt.Errorf("Update HasID") }

	if len(m.Imagen) == 0 {
		m.Imagen = ""
	}

	m.UpdateAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil { return err}

	return nil
}

func (p Producto) UpdateEsceptImage(ID uuid.UUID, updatedProducto model.Producto) error{
	//Obtener el producto por su ID
	m, err := p.GetByID(ID)
	if err != nil { return err }

	//Actualizar todos los campos excepto la imagen
	m.Nombre = updatedProducto.Nombre
	m.Descripcion = updatedProducto.Descripcion
	m.Activo = updatedProducto.Activo
	m.Precio = updatedProducto.Precio
	m.Time = updatedProducto.Time
	m.Calorias = updatedProducto.Calorias

	//Actualizar otros campos segun la estructura de tu modelo
	m.UpdateAt = time.Now().Unix()

	//Actualizar la categoria en el almacenamiento
	err = p.storage.Update(&m)
	if err != nil { return err }

	return nil
}

func (p Producto) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil { return err }

	return nil
}

func (p Producto) GetByID(ID uuid.UUID) (model.Producto, error) {
	producto, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Producto{}, fmt.Errorf("producto: %w", err)
	}

	return producto, nil
}

func (p Producto) GetByCategoryID(idCategoria uuid.UUID) (model.Productos, error) {
	productos, err := p.storage.GetByCategoryID(idCategoria)
	if err != nil {
		return model.Productos{}, fmt.Errorf("producto: %w", err)
	}
	return productos, nil
}

func(p Producto) GetAll() (model.Productos, error) {
	productos, err := p.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("producto: %w", err)
	}
	return productos, nil
}

