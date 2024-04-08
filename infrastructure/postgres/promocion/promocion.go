package promocion

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/delivery/infrastructure/postgres"
	"github.com/ninosistemas10/delivery/model"
)

const table = "promocion"
var field = []string{
	"id",
	"nombre",
	"description",
	"image",
	"precio",
	"features",
	"activo",
	"created",
	"updated",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, field)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, field)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, field)
)

type Promocion struct {
	db * pgxpool.Pool
}

func New(db * pgxpool.Pool) Promocion {
	return Promocion{db: db}
}

func(p Promocion) Create(m * model.Promocion) error {
	_, err := p.db.Exec (
		context.Background(),
		psqlInsert,
		m.ID,
		m.Nombre,
		m.Description,
		m.Image,
		m.Precio,
		m.Features,
		m.Activo,
		m.Created,
		postgres.Int64ToNull(m.Updated),
	)
	if err != nil { return err }
	return nil
}

func (p Promocion) Update(m * model.Promocion) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Nombre,
		m.Description,
		m.Image,
		m.Precio,
		m.Features,
		m.Activo,
		m.Updated,
		m.ID,
	)
	if err != nil { return err }
	return nil
}

func (p Promocion) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil { return err }
	return nil
}

func (p Promocion) GetByID(ID uuid.UUID) (model.Promocion, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return p.scanRow(row)
}

func (p Promocion) GetAll() (model.Promociones, error) {
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)

	if err != nil { return nil, err }
	defer rows.Close()

	var ms model.Promociones
	for rows.Next(){
		m, err := p.scanRow(rows)
		if err != nil { return nil, err }
		ms = append(ms, m)
	}
	return ms, nil
}

func (p Promocion) scanRow(s pgx.Row) (model.Promocion, error) {
	m := model.Promocion{}
	updatedNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Description,
		&m.Image,
		&m.Precio,
		&m.Features,
		&m.Activo,
		&m.Created,
		&updatedNull,
	)
	if err != nil { return m, err }

	m.Updated = updatedNull.Int64
	return m, nil
}
