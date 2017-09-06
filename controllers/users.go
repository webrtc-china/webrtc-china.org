package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webrtc-china.org/middlewares"
	"webrtc-china.org/models"
	"webrtc-china.org/views"
)

type usersImpl struct{}

func RegisterUsers(router *gin.Engine) {
	impl := &usersImpl{}
	router.POST("/account/signup", impl.signUp)
	router.POST("/account/signin", impl.signIn)
	router.GET("/u/:name", impl.show)
}

type userRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func (impl *usersImpl) signUp(c *gin.Context) {
	var body userRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if user, e := models.CreateUser(middlewares.Context(c), body.Username, body.FullName, body.Email, body.Password); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		middlewares.SetupCookie(c.Writer, c.Request, user)
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}

func (impl *usersImpl) signIn(c *gin.Context) {
	var body userRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if user, e := models.AuthUserWithPassword(middlewares.Context(c), body.Username, body.Email, body.Password); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		middlewares.SetupCookie(c.Writer, c.Request, user)
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}

func (impl *usersImpl) show(c *gin.Context) {
	name := c.Params.ByName("name")
	ctx := middlewares.Context(c)
	if user, err := models.FindUserByName(ctx, name); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}
