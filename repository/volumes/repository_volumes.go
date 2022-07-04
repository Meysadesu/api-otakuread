package volumes

import "github.com/Meysadesu/otakuread/model"

type RepositoryVolumes interface {
	Create(volume model.Volume) error
	FindAll() ([]model.Volume, error)
	FindByID(id string) (model.Volume, error)
}
