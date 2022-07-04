package auth

import (
	"net/http"

	"github.com/Meysadesu/otakuread/config/http/validate"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/service/auth"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service auth.ServiceAuth
}

func NewHandlerAuth(service auth.ServiceAuth) *handler {
	return &handler{service: service}

}

func (h *handler) Register(c *fiber.Ctx) error {
	bodyAuth := new(model.Auth)

	err := c.BodyParser(bodyAuth)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot parse body header!!",
			Data:     nil,
		})
	}

	errors := validate.ValidateStruct(bodyAuth)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(entities.WebResponse{
			Code:     http.StatusBadRequest,
			Messages: &errors,
			Data:     nil,
		})
	}

	response := h.service.Register(*bodyAuth)
	return c.Status(response.Code).JSON(response)
}

func (h *handler) Login(c *fiber.Ctx) error {
	body := new(model.Auth)
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot body parse header!!",
			Data:     nil,
		})
	}

	errors := validate.ValidateStruct(body)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(entities.WebResponse{
			Code:     http.StatusBadRequest,
			Messages: &errors,
			Data:     nil,
		})
	}

	response := h.service.Login(*body)
	return c.Status(response.Code).JSON(response)
}
