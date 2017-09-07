package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"webrtc-china.org/models"
)

type contextValueKey struct{ int }

var keyCurrentUser = contextValueKey{1000}

func SetupCookie(w http.ResponseWriter, r *http.Request, user *models.User) {
	cookie := user.Authentication(r.Context())
	cookie.Domain = FetchCookieDomain(r)
	http.SetCookie(w, cookie)
}

func FetchCookieDomain(r *http.Request) string {
	domain := strings.Split(r.Host, ":")[0]
	if len(strings.Split(domain, ".")) > 2 {
		domain = domain[strings.Index(domain, ".")+1:]
	}
	return domain
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		cook, err := c.Request.Cookie("Authorization")
		if err != nil {
			log.Println(err)
			// unauthentication
			return
		}
		user, err := models.AuthenticateUserWithToken(Context(c), cook.Value)
		if err != nil {
			// TODO
		} else if user != nil {
			SetupCookie(c.Writer, c.Request, user)
			WithUser(c, user)
		}
		c.Next()
	}
}
