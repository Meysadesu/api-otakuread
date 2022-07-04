package novels

import "github.com/Meysadesu/otakuread/model"

type RepositoryNovels interface {
	Create(novel model.Novel) error
	Find() ([]model.Novel, error)
	FindByTitle(title string) (model.Novel, error)
	Delete(id string) error
	Update(novel model.Novel) error
}
