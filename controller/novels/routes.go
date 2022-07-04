package novels

import (
	"github.com/Meysadesu/otakuread/config/database"
	"github.com/Meysadesu/otakuread/config/http/middleware/jwt"
	"github.com/Meysadesu/otakuread/repository/novels"
	srv "github.com/Meysadesu/otakuread/service/novels"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	db, err := database.MysqlConnect()
	if err != nil {
		panic(err.Error())
	}
	repo := novels.NewRepositoryNovel(db)
	serv := srv.NewServiceNovels(repo)
	handler := NewHandler(serv)

	novel := app.Group("/v2")
	novel.Post("/novels", jwt.JwtMiddleware(), handler.Create)
	app.Get("/v2/novels", handler.Find)
	app.Get("/v2/novel", handler.FindByTitle)
	novel.Put("/novels", handler.Update)
	novel.Delete("/novels", handler.Delete)
}
