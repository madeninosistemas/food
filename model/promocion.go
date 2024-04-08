package model

import (
	"encoding/json"

	"github.com/google/uuid"
)


type Promocion struct {
	ID 				uuid.UUID 		`json:"id"`
	Nombre			string			`json:"nombre"`
	Description		string			`json:"description"`
	Image			string 			`json:"image"`
	Precio			float64			 `json:"precio"`
	Features		json.RawMessage	`json:"features"`
	Activo			bool			`json:"activo"`
	Created			int64			`json:"created"`
	Updated			int64			`json:"updated"`
}

func (p Promocion) HasID() bool {
	return p.ID != uuid.Nil
}

type Promociones []Promocion

func (p Promociones) IdEmpty() bool {
	return len(p) == 0
}

