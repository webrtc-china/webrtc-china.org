package middlewares

import (
	"context"

	"webrtc-china.org/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"webrtc-china.org/session"
)

const keyContext = "context"
const keyUser = "user"

func WithContext(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = session.WithDatabase(ctx, db)
		c.Set(keyContext, ctx)
		c.Next()
	}
}

func WithUser(c *gin.Context, user *models.User) {
	c.Set(keyUser, user)
}

func Context(c *gin.Context) context.Context {
	return c.MustGet(keyContext).(context.Context)
}

func User(c *gin.Context) *models.User {
	return c.MustGet(keyUser).(*models.User)
}
