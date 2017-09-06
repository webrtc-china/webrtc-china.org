package middlewares

import (
	"net/http"
	"strings"

	"webrtc-china.org/models"
)

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
