package server

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"jjaa.me/http/controllers"
)

func SetRoutes(router *gin.Engine) {

	router.GET("/", controllers.WelcomeIndex)

	sessions := router.Group("/sessions")
	sessions.GET("/new", controllers.SessionsNew)
	sessions.POST("/", controllers.SessionsCreate)
	sessions.POST("/destroy", controllers.SessionsDestroy)

	videos := router.Group("/videos")
	videos.GET("/new", controllers.VideosNew)
	videos.POST("/", controllers.VideosCreate)
	videos.GET("/", controllers.VideosIndex)
	videos.GET("/all", controllers.VideosAllIndex)
	videos.GET("/upload/:name", controllers.VideosUpload)
	videos.POST("/destroy", controllers.VideosDestroy)
	videos.POST("/file/:name", controllers.VideosFile)
	videos.GET("/view/:name", controllers.VideosShow)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func RunHttpAndHttps(router *gin.Engine) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("jjaa.me"),
		Cache:      autocert.DirCache("/certs"),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: router,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	server.ListenAndServeTLS("", "")
}
