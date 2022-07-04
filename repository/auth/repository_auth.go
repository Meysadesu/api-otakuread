package auth

import "github.com/Meysadesu/otakuread/model"

type RepositoryAuth interface {
	Create(auth model.Auth) error
	FindByUsername(username string) (model.Auth, error)
	Count(username string) int64
}
