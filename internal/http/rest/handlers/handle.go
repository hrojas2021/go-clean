package handlers

import "github.com/hugo.rojas/custom-api/internal/iface"

func New(srv iface.Service) Handle {
	return Handle{srv}
}

type Handle struct {
	service iface.Service
	// resp    iface.Response
}
