package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/http/rest/middlewares"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func createJWT(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service, conf *conf.Configuration) *bunrouter.CompatRouter {
	// https://bunrouter.uptrace.dev/guide/golang-router.html#installation
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		// add default error routes
	).Compat()

	h := handlers.New(service)

	r.GET("/hugo", func(w http.ResponseWriter, req *http.Request) { // CONVERT INTO SERVICE
		token, err := createJWT(conf.JWT.SECRET)
		if err != nil {
			return
		}
		fmt.Fprint(w, token)
	})

	api := r.NewGroup("/api", bunrouter.Use(middlewares.Authenticate))
	api.WithGroup("/", func(g *bunrouter.CompatGroup) {
		g.GET("/users", h.ListUsers)
		g.POST("/rooms", h.SaveRoom)
	})

	return r
}
