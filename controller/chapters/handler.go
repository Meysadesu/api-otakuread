package chapters

import (
	"net/http"

	"github.com/Meysadesu/otakuread/config/http/validate"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/service/chapters"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service chapters.ServiceChapters
}

func NewHandler(service chapters.ServiceChapters) *handler {
	return &handler{service: service}
}

func (h *handler) Create(c *fiber.Ctx) error {
	body := new(model.Chapter)

	err := c.BodyParser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "error parse body",
			Data:     nil,
		})
	}

	errors := validate.ValidateStruct(body)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(entities.WebResponse{
			Code:     http.StatusBadRequest,
			Messages: errors,
			Data:     nil,
		})
	}

	response := h.service.Create(*body)
	return c.Status(response.Code).JSON(response)
}

func (h *handler) FindByCH(c *fiber.Ctx) error {
	ch := c.Query("ch")
	response := h.service.FindByCH(ch)
	return c.Status(response.Code).JSON(response)
}
