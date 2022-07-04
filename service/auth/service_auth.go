package auth

import (
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
)

type ServiceAuth interface {
	Register(auth model.Auth) entities.WebResponse
	Login(auth model.Auth) entities.WebResponse
}
