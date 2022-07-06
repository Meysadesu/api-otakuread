package chapters

import (
	"net/http"

	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/repository/chapters"
	"gorm.io/gorm"
)

type service struct {
	repository chapters.RepositoryChapters
}

func NewServiceChapters(repository chapters.RepositoryChapters) ServiceChapters {
	return &service{repository: repository}
}

func (s *service) Create(chapters model.Chapter) entities.WebResponse {
	err := s.repository.Create(chapters)
	if err != nil {
		return entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot create data",
			Data:     nil,
		}
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "create data successfully",
		Data:     nil,
	}
}

func (s *service) FindByCH(ch string) entities.WebResponse {
	response, err := s.repository.FindByCH(ch)
	if err == gorm.ErrRecordNotFound {
		return entities.WebResponse{
			Code:     http.StatusNotFound,
			Messages: "data not found!",
			Data:     nil,
		}
	}

	resp := model.ChapterResponse{
		ID:       response.ID,
		IDVolume: response.ID,
		Ch:       response.Ch,
		Header:   response.Header,
		Ctx:      response.Ctx,
	}

	return entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "successfully",
		Data:     resp,
	}
}
