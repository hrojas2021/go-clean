package infrastructure

import (
	"github.com/jmoiron/sqlx"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	DB      *sqlx.DB
	Config  *conf.Configuration
	Handler *httprouter.Router
}

func NewAPI(db *sqlx.DB, config *conf.Configuration) *API {
	return &API{
		DB:     db,
		Config: config,
	}
}
