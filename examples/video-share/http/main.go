package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"jjaa.me/http/controllers"
	"jjaa.me/http/server"
	"jjaa.me/persist"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	controllers.Db = persist.Connection()

	router := gin.Default()
	server.SetRoutes(router)
	if os.Getenv("GIN_MODE") == "release" {
		server.AddTemplates(router, "/")
		router.Static("/assets", "/assets")
		server.RunHttpAndHttps(router)
	} else {
		server.AddTemplates(router, "./")
		router.Static("/assets", "assets")
		router.Run(":3000")
	}
}
