package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"jjaa.me/http/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()
	server.SetRoutes(router)
	if os.Getenv("GIN_MODE") == "release" {
		path := "/root/MCLOBwVMs/examples/video-share/http/"
		server.AddTemplates(router, path)
		router.Static("/assets", path+"assets")
		server.RunHttpAndHttps(router)
	} else {
		server.AddTemplates(router, "./")
		router.Static("/assets", "assets")
		router.Run(":3000")
	}
}
