package volumes

import (
	"net/http"

	"github.com/Meysadesu/otakuread/config/http/validate"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/service/volumes"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service volumes.ServiceVolumes
}

func NewHandler(service volumes.ServiceVolumes) *handler {
	return &handler{service: service}
}

func (h *handler) Create(c *fiber.Ctx) error {
	body := new(model.Volume)
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: err,
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

func (h *handler) FindAll(c *fiber.Ctx) error {
	response := h.service.Find()
	return c.Status(response.Code).JSON(response)
}

func (h *handler) FindByID(c *fiber.Ctx) error {
	id := c.Query("id")
	response := h.service.FindByID(id)
	return c.Status(response.Code).JSON(response)
}
