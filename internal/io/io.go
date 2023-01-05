package io

import (
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/hugo.rojas/custom-api/internal/io/database"
)

type IO struct {
	*database.Database
}

func New(d *database.Database) iface.IO {
	return &IO{
		Database: d,
	}
}
