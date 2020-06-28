package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	router := gin.Default()

	//router.Static("/assets", prefix+"assets")
	//router.GET("/", controllers.WelcomeIndex)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	go router.Run(fmt.Sprintf(":%s", port))

	for {
		time.Sleep(time.Second)
	}

}
