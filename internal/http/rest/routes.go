package rest

import (
	"io"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service) *bunrouter.CompatRouter {
	// https://bunrouter.uptrace.dev/guide/golang-router.html#installation
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		// add default error routes
	).Compat()

	// r.Use(Telemetry())

	h := handlers.New(service)
	// 	// r.Use(Recoverer(resp))  CHECK Recoverer and Check if Resp is needed

	api := r.NewGroup("/api", bunrouter.Use(errorHandler))
	api.WithGroup("/", func(g *bunrouter.CompatGroup) {
		g.GET("/campaigns/:id", h.GetCampaign)
	})

	return r
}

func errorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {

	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)

		switch err := err.(type) {
		case nil:
			// no error
		case HTTPError:
			w.WriteHeader(err.statusCode)
			_ = bunrouter.JSON(w, err)
		default:
			httpErr := NewHTTPError(err)
			w.WriteHeader(httpErr.statusCode)
			_ = bunrouter.JSON(w, httpErr)
		}

		return err
	}
}

type HTTPError struct {
	statusCode int

	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(err error) HTTPError {
	switch err {
	case io.EOF:
		return HTTPError{
			statusCode: http.StatusBadRequest,

			Code:    "eof",
			Message: "EOF reading HTTP request body",
		}
	}

	return HTTPError{
		statusCode: http.StatusInternalServerError,

		Code:    "internal",
		Message: "Internal server error",
	}
}
