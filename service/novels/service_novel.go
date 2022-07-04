package novels

import (
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
)

type ServiceNovels interface {
	Find() entities.WebResponse
	Create(novels model.Novel) entities.WebResponse
	FindByTitle(title string) entities.WebResponse
	Update(novel model.Novel) entities.WebResponse
	Delete(id string) entities.WebResponse
}
