package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"webrtc-china.org/middlewares"
	"webrtc-china.org/models"
	"webrtc-china.org/views"
)

type repliesImpl struct{}

type replyRequest struct {
	Content string `json:"content" binding:"required"`
}

func RegisterReplies(router *gin.Engine) {
	impl := &repliesImpl{}
	router.PUT("/topics/:topic_id/replies", impl.create)
}

func (impl *repliesImpl) create(c *gin.Context) {
	var bodyRequest replyRequest
	user := middlewares.User(c)
	topicId, _ := strconv.ParseInt(c.Params.ByName("topic_id"), 10, 64)
	if err := c.BindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if rp, e := models.CreateReply(middlewares.Context(c), bodyRequest.Content, user.Id, topicId); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildReplyView(topicId, rp, user))
	}
}
