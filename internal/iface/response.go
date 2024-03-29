//go:generate go run github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=mock/response.go
package iface

import "net/http"

type Response interface {
	Fail(w http.ResponseWriter, r *http.Request, err error)
	Failf(w http.ResponseWriter, r *http.Request, format string, a ...interface{})
	JSON(w http.ResponseWriter, r *http.Request, data interface{})
}
