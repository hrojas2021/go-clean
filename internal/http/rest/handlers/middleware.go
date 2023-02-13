package handlers

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/uptrace/bunrouter"
)

const prefixLen = len("Bearer ")

func (h *Handle) Authenticate(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var rawJWT string
		if raw := req.Header.Get("Authorization"); len(raw) > prefixLen {
			rawJWT = raw[prefixLen:]

			token, err := jwt.Parse(rawJWT, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized")) //nolint:errcheck // error response
				}
				return h.service.GetSecret(), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error())) //nolint:errcheck // error response
			}

			if token.Valid {
				return next(w, req.WithContext(req.Context()))
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("not authorized")
	}
}

func CorsMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			return next(w, req)
		}

		h := w.Header()

		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Credentials", "true")

		// CORS preflight.
		if req.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
			h.Set("Access-Control-Allow-Headers", "authorization,content-type")
			h.Set("Access-Control-Max-Age", "86400")
			return nil
		}

		return next(w, req)
	}
}

// func Telemetry(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
// 	return func(w http.ResponseWriter, req bunrouter.Request) error {
// 		spanHandler := func(w http.ResponseWriter, r *http.Request) {
// 			traceID := trace.SpanFromContext(r.Context()).SpanContext().TraceID().String()

// 			// add trace id to the http response
// 			w.Header().Add("Trace-ID", traceID)

// 			next.ServeHTTP(w, r)
// 		}
// 		return otelhttp.NewHandler(http.HandlerFunc(spanHandler), "request")
// 	}
// }
