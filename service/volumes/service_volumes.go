package volumes

import (
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
)

type ServiceVolumes interface {
	Find() entities.WebResponse
	FindByID(id string) entities.WebResponse
	Create(volume model.Volume) entities.WebResponse
}
