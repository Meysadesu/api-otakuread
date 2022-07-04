package novels

import (
	"fmt"
	"net/http"

	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/repository/novels"
	"gorm.io/gorm"
)

type service struct {
	repository novels.RepositoryNovels
}

func NewServiceNovels(repository novels.RepositoryNovels) ServiceNovels {
	return &service{repository: repository}
}

func (s *service) Create(novels model.Novel) entities.WebResponse {
	err := s.repository.Create(novels)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot insert data!!",
			Data:     nil,
		}
	}

	response, err := s.repository.FindByTitle(novels.Title)
	if err == gorm.ErrRecordNotFound {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "data not found!",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     response,
	}
}

func (s *service) Find() entities.WebResponse {
	dataResponse, err := s.repository.Find()
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot query data!!",
			Data:     nil,
		}
	}

	if len(dataResponse) == 0 {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "cannot request, data null",
			Data:     dataResponse,
		}
	}

	var response []model.NovelResponse
	for _, v := range dataResponse {
		response = append(response, model.NovelResponse{
			ID:          v.ID,
			Title:       v.Title,
			Cover:       v.Title,
			Type:        v.Type,
			Author:      v.Author,
			Status:      v.Status,
			Ranting:     v.Ranting,
			Genre:       v.Genre,
			Description: v.Description,
		})
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "succesfully",
		Data:     response,
	}
}

func (s *service) FindByTitle(title string) entities.WebResponse {
	response, err := s.repository.FindByTitle(title)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot query data!!",
			Data:     nil,
		}
	}

	if err == gorm.ErrRecordNotFound {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "data not found!",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     response,
	}
}

func (s *service) Delete(id string) entities.WebResponse {
	err := s.repository.Delete(id)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannnot delete data!!",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: fmt.Sprintf("delete data id %s success", id),
		Data:     nil,
	}
}

func (s *service) Update(novel model.Novel) entities.WebResponse {
	err := s.repository.Update(novel)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot update data!",
			Data:     nil,
		}
	}

	response, err := s.repository.FindByTitle(novel.Title)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot query data!",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: fmt.Sprintf("update data id %d success", novel.ID),
		Data:     response,
	}
}
