package integration

import (
	"context"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/service"
)

var (
	fixture iface.Service
	token   string
	ctx     = context.Background()
)

func init() {
	cf := conf.LoadViperConfig()
	db := database.InitDB(cf)
	io := io.New(database.New(db))
	fixture = service.New(cf, io)
	token = getToken()
}

func createGenericRoom(name string) (*models.Room, error) {
	r := &models.Room{
		Name: name,
	}

	err := fixture.SaveRoom(ctx, r)
	return r, err
}

func getToken() string {
	user := models.User{
		Username: "hrojas",
		Password: "12345",
	}
	t, _ := fixture.Login(ctx, user)
	return t.Token
}
