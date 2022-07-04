package novels

import (
	"net/http"

	"github.com/Meysadesu/otakuread/config/http/validate"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/Meysadesu/otakuread/model"
	"github.com/Meysadesu/otakuread/service/novels"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service novels.ServiceNovels
}

func NewHandler(service novels.ServiceNovels) *handler {
	return &handler{service: service}
}

func (h *handler) Create(c *fiber.Ctx) error {
	body := new(model.Novel)

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
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	response := h.service.Create(*body)
	return c.Status(response.Code).JSON(response)
}

func (h *handler) Find(c *fiber.Ctx) error {
	response := h.service.Find()
	return c.Status(response.Code).JSON(response)
}

func (h *handler) FindByTitle(c *fiber.Ctx) error {
	title := c.Query("title")
	response := h.service.FindByTitle(title)
	return c.Status(response.Code).JSON(response)
}

func (h *handler) Update(c *fiber.Ctx) error {
	body := new(model.Novel)

	err := c.BodyParser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(entities.WebResponse{
			Code:     http.StatusInternalServerError,
			Messages: "cannot body parse!!",
			Data:     nil,
		})
	}

	response := h.service.Update(*body)
	return c.Status(response.Code).JSON(response)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id := c.Query("id")
	response := h.service.Delete(id)
	return c.Status(response.Code).JSON(response)

}
