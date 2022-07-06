package volumes

import (
	"net/http"

	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/repository/volumes"
	"gorm.io/gorm"
)

type service struct {
	repository volumes.RepositoryVolumes
}

func NewServiceNovels(repository volumes.RepositoryVolumes) ServiceVolumes {
	return &service{repository: repository}
}

func (s *service) Create(volume model.Volume) entities.WebResponse {
	err := s.repository.Create(volume)
	if err != err {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot create data",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     nil,
	}
}

func (s *service) FindByID(id string) entities.WebResponse {
	response, err := s.repository.FindByID(id)
	if err == gorm.ErrRecordNotFound {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "data not found!",
			Data:     nil,
		}
	}

	resp := model.VolumeResponse{
		ID:          response.ID,
		TitleNovel:  response.TitleNovel,
		WhatsVolume: response.WhatsVolume,
		Chapter:     response.Chapter,
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     resp,
	}
}

func (s *service) Find() entities.WebResponse {
	response, err := s.repository.FindAll()
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: err,
			Data:     nil,
		}
	}

	if len(response) == 0 {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "data not found",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     nil,
	}
}
