package controllers

import (
	"net/http"

	"github.com/go-pg/pg"

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
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (impl *usersImpl) signUp(c *gin.Context) {
	var body userRequest
	db := c.MustGet(middlewares.KeyDatabase).(*pg.DB)
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if user, e := models.CreateUser(db, body.Username, body.FullName, body.Email, body.Password); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}

func (impl *usersImpl) signIn(c *gin.Context) {
	var body userRequest
	db := c.MustGet(middlewares.KeyDatabase).(*pg.DB)
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if user, e := models.AuthUserWithPassword(db, body.Username, body.Email, body.Password); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}

func (impl *usersImpl) show(c *gin.Context) {
	name := c.Params.ByName("name")
	db := c.MustGet(middlewares.KeyDatabase).(*pg.DB)
	if user, err := models.FindUserByName(db, name); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, views.BuildUserView(user))
	}
}
