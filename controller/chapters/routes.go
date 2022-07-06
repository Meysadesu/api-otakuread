package chapters

import (
	"github.com/Meysadesu/otakuread/config/database"
	"github.com/Meysadesu/otakuread/config/http/middleware/jwt"
	"github.com/Meysadesu/otakuread/repository/chapters"
	srv "github.com/Meysadesu/otakuread/service/chapters"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	db, _ := database.MysqlConnect()

	repo := chapters.NewRepositoryChapter(db)
	serv := srv.NewServiceChapters(repo)
	handler := NewHandler(serv)

	chapter := app.Group("/v2")
	app.Post("/v2/chapters", jwt.JwtMiddleware(), handler.Create)
	chapter.Get("/chapters", handler.FindByCH)
}
