package main

import (
	"log"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"webrtc-china.org/controllers"
	"webrtc-china.org/middlewares"
	"webrtc-china.org/models"
)

func main() {
	db := models.InitDb()

	router := gin.Default()
	router.GET("/", index)
	router.Use(middlewares.Context(db))
	controllers.RegisterUsers(router)
	controllers.RegisterTopics(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	gracehttp.Serve(server)
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "hello world"})
	log.Println("hello index")
}
