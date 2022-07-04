package chapters

import "github.com/Meysadesu/otakuread/model"

type RepositoryChapters interface {
	Create(chapters model.Chapter) error
	FindByCH(chapter string) (model.Chapter, error)
}
