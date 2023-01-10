package middlewares

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/uptrace/bunrouter"
)

const prefixLen = len("Bearer ")

var SECRET = []byte("ZTP02X517M4PUND7") // FIX THIS

func Authenticate(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {

	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var rawJWT string
		if raw := req.Header.Get("Authorization"); len(raw) > prefixLen {

			rawJWT = raw[prefixLen:]

			token, err := jwt.Parse(rawJWT, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				return next(w, req.WithContext(req.Context()))
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("not authorized")
	}
}
