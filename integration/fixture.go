package integration

import (
	"context"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/service"
)

var (
	fixture iface.Service
	ctx     = context.Background()
)

func init() {
	cf := conf.LoadViperConfig()
	db := database.InitDB(cf)
	io := io.New(database.New(db))
	fixture = service.New(cf, io)
}
