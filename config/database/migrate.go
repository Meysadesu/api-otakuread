package database

import (
	"github.com/Meysadesu/otakuread/model"
)

type Migrate struct {
	model interface{}
}

func MigrateTable() []Migrate {
	return []Migrate{
		Migrate{model: model.Novel{}},
		Migrate{model: model.Volume{}},
		Migrate{model: model.Chapter{}},
		Migrate{model: model.Auth{}},
	}
}
