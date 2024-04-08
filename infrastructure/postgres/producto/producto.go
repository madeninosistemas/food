package producto

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/delivery/infrastructure/postgres"
	"github.com/ninosistemas10/delivery/model"
)


const table = "productos"

var fields = [] string{
	"id",
	"idcategoria",
	"nombre",
	"precio",
	"imagen",
	"descripcion",
	"activo",
	"time",
	"calorias",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
	psqlGetAllByCategory = postgres.BuilddSQLSelectByCategory(table, fields)

)

type Producto struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Producto {
	return Producto{db}
}

func (p Producto) Create(m *model.Producto) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.IdCategoria,
		m.Nombre,
		m.Precio,
		m.Imagen,
		m.Descripcion,
		m.Activo,
		m.Time,
		m.Calorias,
		m.CreateAt,
		postgres.Int64ToNull(m.UpdateAt),
	)
	if err != nil { return err }
	return nil
}

func (p Producto) Update(m * model.Producto) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.IdCategoria,
		m.Nombre,
		m.Precio,
		m.Imagen,
		m.Descripcion,
		m.Activo,
		m.Time,
		m.Calorias,
		m.UpdateAt,
		m.ID,
	)

	if err != nil { return err }
	return nil
}

func (p Producto) UpdateEsceptImage(ID uuid.UUID, updatedProducto model.Producto) error {
	//Actualizar todos los campos excepto la imagen
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		updatedProducto.Nombre,
		updatedProducto.Descripcion,
		updatedProducto.Activo,
		updatedProducto.Precio,
		updatedProducto.Time,
		updatedProducto.Calorias,

	)
	if err != nil { return  err}
	return nil
}


func(p Producto) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if  err != nil { return err }
	return nil
}

func (p Producto) GetByID(ID uuid.UUID) (model.Producto, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return p.scanRow(row)
}

func (p Producto) GetByCategoryID(idCategoria uuid.UUID) (model.Productos, error) {
	query := psqlGetAllByCategory
	rows, err := p.db.Query(
		context.Background(),
		query + " WHERE idcategoria = $1",
		idCategoria,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos model.Productos
	for rows.Next() {
		producto, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}

	return productos, nil
}





func (p Producto) GetAll() (model.Productos, error){
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil { return nil, err }

	defer rows.Close()

	var ms model.Productos
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil { return nil, err }

		ms = append(ms, m)
	}

	return ms, nil
}

func (p Producto) scanRow(s pgx.Row) (model.Producto, error) {
	m := model.Producto{}

	updateAtNull := sql.NullInt64{}
	err := s.Scan(
		&m.ID,
		&m.IdCategoria,
		&m.Nombre,
		&m.Precio,
		&m.Imagen,
		&m.Descripcion,
		&m.Activo,
		&m.Time,
		&m.Calorias,
		&m.CreateAt,
		&updateAtNull,
	)
	if err != nil { return m, err }

	m.UpdateAt = updateAtNull.Int64

	return m, nil
}






