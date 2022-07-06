package server

import (
	"net/http"

	"github.com/Meysadesu/otakuread/controller/auth"
	"github.com/Meysadesu/otakuread/controller/chapters"
	"github.com/Meysadesu/otakuread/controller/novels"
	"github.com/Meysadesu/otakuread/controller/volumes"
	"github.com/Meysadesu/otakuread/entities"
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	app := fiber.New()
	app.Get("/", HomeHandler)
	auth.Routes(app)
	novels.Routes(app)
	chapters.Routes(app)
	volumes.Routes(app)
	return app
}

func HomeHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(entities.WebResponse{
		Code:     http.StatusOK,
		Messages: "Welcome to Otakuread Rest Api - Hallo Dunia",
		Data:     nil,
	})
}
