package server

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func SetRoutes(router *gin.Engine) {

	//router.Static("/assets", prefix+"assets")
	//router.GET("/", controllers.WelcomeIndex)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome")
	})
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
