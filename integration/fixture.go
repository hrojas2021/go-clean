package integration

import (
	"context"
	"fmt"

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
	fixtures.srv = service.New(cf, io)
	token = fixtures.getToken()
	httpClient = getHTTPClient(token)
}

func (f *Fixtures) createGenericRoom(name string) (*models.Room, error) {
	r := &models.Room{
		Name: name,
	}

	err := f.srv.SaveRoom(ctx, r)
	return r, err
}

func (f *Fixtures) getToken() string {
	user := models.User{
		Username: "hrojas",
		Password: "12345",
	}
	t, _ := f.srv.Login(ctx, user)
	return t.Token
}
