/*
@ Restful-Api remake
@ name : Otakuread-Restful-Api
@ description : Restful-Api for Website Otakuread
*/

package main

import (
	"os"

	"github.com/Meysadesu/otakuread/config/http/server"
)

func main() {
	app := server.Server()
	app.Listen(":" + os.Getenv("PORT"))
}
