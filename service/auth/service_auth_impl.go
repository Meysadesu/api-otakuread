package auth

import (
	"net/http"

	"github.com/Meysadesu/otakuread/config/http/middleware/jwt"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/repository/auth"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository auth.RepositoryAuth
}

func NewServiceAuth(repository auth.RepositoryAuth) ServiceAuth {
	return &service{repository: repository}
}

func (s *service) Register(auth model.Auth) entities.WebResponse {

	userExist := s.repository.Count(auth.Username)
	if userExist == 1 {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "user already exist!",
			Data:     nil,
		}
	}

	if auth.Username == "" || auth.Email == "" {
		return entities.WebResponse{
			Code:     http.StatusBadRequest,
			Messages: "invalid body!!",
			Data:     nil,
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "register failed, cannot generate password",
			Data:     nil,
		}
	}

	err = s.repository.Create(model.Auth{
		Username: auth.Username,
		Email:    auth.Email,
		Password: string(password),
		Role:     auth.Role,
	})

	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: err,
			Data:     nil,
		}
	}

	userData, _ := s.repository.FindByUsername(auth.Username)
	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "success",
		Data:     userData,
	}
}

func (s *service) Login(auth model.Auth) entities.WebResponse {
	data, err := s.repository.FindByUsername(auth.Username)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "user not found!",
			Data:     nil,
		}
	}

	password := data.Password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(auth.Password))
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusBadRequest,
			Messages: "password wrong!!",
			Data:     nil,
		}
	}

	if data.Role != "admin" {
		response := model.UserResponse{
			ID:       data.ID,
			Username: data.Username,
			Email:    data.Email,
			Password: data.Password,
		}

		return entities.WebResponse{
			Code:     http.StatusOK,
			Messages: "Welcome",
			Data:     response,
		}
	}

	token, err := jwt.CreateToken(data.Username)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot generate token!!",
			Data:     nil,
		}
	}

	response := model.AdminResponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Token:    token,
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "success",
		Data:     response,
	}
}
