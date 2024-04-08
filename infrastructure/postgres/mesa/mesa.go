package mesa

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/delivery/infrastructure/postgres"
	"github.com/ninosistemas10/delivery/model"
)

const table = "mesa"

var fields = []string{
	"id",
	"nombre",
	"url",
	"images",
	"activo",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type Mesa struct {
	db *pgxpool.Pool
}

func New(db * pgxpool.Pool) Mesa {
	return Mesa{db}
}

func (me Mesa) Create(m * model.Mesa) error {
	_, err := me.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Nombre,
		m.Url,
		m.Images,
		m.Activo,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil { return err }
	return nil
}

func (me Mesa) Update(m * model.Mesa) error {
	_, err := me.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Nombre,
		m.Url,
		m.Images,
		m.Activo,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil { return err }
	return nil
}

func (me Mesa) Delete(ID uuid.UUID) error {
	_, err := me.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil { return err }
	return nil
}

func (me Mesa) GetByID(ID uuid.UUID) (model.Mesa, error){
	query := psqlGetAll + "WHERE id = $1"
	row := me.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return me.scanRow(row)
}

func (me Mesa) GetAll()(model.Mesas, error) {
	rows, err := me.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil { return nil, err }
	defer rows.Close()

	var ms model.Mesas
	for rows.Next() {
		m, err := me.scanRow(rows)
		if err != nil { return nil, err }
		ms = append(ms, m)
	}

	return ms, nil
}


func (me Mesa) scanRow(s pgx.Row) (model.Mesa, error) {
	m := model.Mesa{}
	updateAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Url,
		&m.Images,
		&m.Activo,
		&m.CreatedAt,
		&updateAtNull,
	)

	if err != nil { return m, err }

	m.UpdatedAt = updateAtNull.Int64
	return m, nil
}
