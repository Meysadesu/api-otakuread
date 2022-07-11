/*
@ Restful-Api remake
@ name : Otakuread-Restful-Api
@ description : Restful-Api for Website Otakuread
*/

package main

import (
	"os"

	"github.com/Meysadesu/otakuread/config/http/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.example")
	app := server.Server()
	app.Listen(":" + os.Getenv("PORT"))
}
