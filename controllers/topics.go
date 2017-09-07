package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webrtc-china.org/middlewares"
	"webrtc-china.org/models"
	"webrtc-china.org/views"
)

type topicsIml struct{}

type topicRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Node    string `json:"node" binding:"required"`
}

func RegisterTopics(router *gin.Engine) {
	impl := &topicsIml{}
	router.PUT("/topics", impl.create)
	router.GET("/topics", impl.topics)
	router.GET("/topics/:id", impl.show)
	router.POST("/topics/:id", impl.update)
}

func (impl *topicsIml) create(c *gin.Context) {
	var bodyRequest topicRequest
	user := middlewares.User(c)
	if err := c.BindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if tp, e := models.CreateTopic(middlewares.Context(c), user, bodyRequest.Title, bodyRequest.Content, bodyRequest.Node); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildTopicView(tp, user))
	}
}

func (impl *topicsIml) topics(c *gin.Context) {

}

func (impl *topicsIml) update(c *gin.Context) {

}

func (impl *topicsIml) show(c *gin.Context) {

}
