package chapters

import (
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
)

type ServiceChapters interface {
	Create(chapters model.Chapter) entities.WebResponse
	FindByCH(ch string) entities.WebResponse
}
