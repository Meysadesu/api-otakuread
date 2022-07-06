package volumes

import (
	"github.com/Meysadesu/otakuread/config/database"
	"github.com/Meysadesu/otakuread/config/http/middleware/jwt"
	"github.com/Meysadesu/otakuread/repository/volumes"
	svc "github.com/Meysadesu/otakuread/service/volumes"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	db, _ := database.MysqlConnect()

	repo := volumes.NewRespositoryVolumes(db)
	srv := svc.NewServiceNovels(repo)
	handler := NewHandler(srv)

	g := app.Group("/v2")
	g.Post("/volumes", jwt.JwtMiddleware(), handler.Create)
	g.Get("/volumes", handler.FindAll)
	app.Get("/v2/volumes", handler.FindByID)
}
