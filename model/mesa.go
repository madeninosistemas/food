package model

import (
	"github.com/google/uuid"
)


type Mesa struct {
	ID			uuid.UUID		`json:"id"`
	Nombre		string			`json:"nombre"`
	Url			string			`json:"url"`
	Images		string			`json:"images"`
	Activo		bool			`json:"activo"`
	CreatedAt	int64			`json:"created_at"`
	UpdatedAt	int64			`json:"updated_at"`
}

func (me Mesa) HasID() bool {
	return me.ID != uuid.Nil
}

type Mesas []Mesa

func(me Mesas) IsEmpty() bool { return len(me) == 0 }
