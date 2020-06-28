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
	controllers.Host = "jjaa.me"

	router := gin.Default()
	server.SetRoutes(router)
	if os.Getenv("GIN_MODE") == "release" {
		server.AddTemplates(router, "/http/")
		router.Static("/assets", "/http/assets")
		server.RunHttpAndHttps(router)
	} else {
		server.AddTemplates(router, "./")
		router.Static("/assets", "assets")
		controllers.Host = "localhost"
		router.Run(":3000")
	}
}
