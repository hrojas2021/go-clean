package handlers

import "github.com/hugo.rojas/custom-api/internal/iface"

func New(srv iface.Service, resp iface.Response) Handle {
	return Handle{srv, resp}
}

type Handle struct {
	service iface.Service
	resp    iface.Response
}
