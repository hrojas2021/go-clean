package service

import (
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/iface"
)

type Service struct {
	config *conf.Configuration
	io     iface.IO
}

func New(conf *conf.Configuration, io iface.IO) iface.Service {
	return &Service{
		config: conf,
		io:     io,
	}
}
