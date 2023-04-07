package handlers

import (
	"github.com/go-playground/validator"
	"github.com/hugo.rojas/custom-api/internal/iface"
)

func New(srv iface.Service, resp iface.Response, v *validator.Validate) Handle {
	return Handle{srv, resp, v}
}

type Handle struct {
	service  iface.Service
	resp     iface.Response
	validate *validator.Validate
}
