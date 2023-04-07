package integration

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/federicoleon/go-httpclient/gohttp"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/service"
)

type Fixtures struct {
	srv iface.Service
}

var (
	token      string
	httpClient gohttp.Client
	ctx        = context.Background()
	localURL   string
	fixtures   Fixtures
)

func init() {
	cf := conf.LoadViperConfig()
	db := database.InitDB(cf)
	io := io.New(database.New(db))
	localURL = fmt.Sprintf("http://localhost:%d", cf.PORT)
	fmt.Printf("\n\n%+v\n\n", cf)
	fixtures.srv = service.New(cf, io)
	token = fixtures.getToken(cf.JWT.SECRET)
	httpClient = getHTTPClient(token)
}

func (f *Fixtures) createGenericRoom(name string) (*models.Room, error) {
	r := &models.Room{
		Name: name,
	}

	err := f.srv.SaveRoom(ctx, r)
	return r, err
}

func (f *Fixtures) getToken(secret string) string {
	var tokenStr string
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	claims["exp"] = time.Now().Add(time.Duration(5) * time.Minute).Unix()
	claims["username"] = "admin-test"

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	fmt.Println("TOKEN", len(tokenStr))
	return tokenStr
}
