package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

const (
	KeyRequest  string = "0"
	KeyDatabase string = "1"
)

func Context(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyDatabase, db)
		c.Next()
	}
}
