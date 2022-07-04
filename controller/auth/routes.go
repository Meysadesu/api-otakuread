package auth

import (
	"github.com/Meysadesu/otakuread/config/database"
	"github.com/Meysadesu/otakuread/repository/auth"
	srv "github.com/Meysadesu/otakuread/service/auth"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	db, err := database.MysqlConnect()
	if err != nil {
		panic(err.Error())
	}
	repo := auth.NewRepositoryAuth(db)
	serv := srv.NewServiceAuth(repo)
	handler := NewHandlerAuth(serv)

	auth := app.Group("/v2/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}
