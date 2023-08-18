package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openbank-ondc-web/src/pkg/controllers"
)

func main() {
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/ondc-site-verification.html")
	router.GET("/ondc-site-verification.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ondc-site-verification.html", gin.H{})
	})
	route := router.Group("/api/v1")
	{
		onSubscribe := route.Group("/ondc/onboarding")
		onSubscribe.POST("/on_subscribe", controllers.OnSubscribe)
	}

	router.Run(":8080")
}
