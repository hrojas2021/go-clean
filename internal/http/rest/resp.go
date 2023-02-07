package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/errors"
)

type Error struct {
	Code string                  `json:"code"`
	Msg  string                  `json:"msg,omitempty"`
	Args *map[string]interface{} `json:"args,omitempty"`
}

type ErrResponse struct {
	Error struct {
		Codes []Error `json:"codes"`
		Msg   string  `json:"msg,omitempty"`
		File  string  `json:"file,omitempty"`
	} `json:"error"`
}

type DefaultResp struct{}

// Fail writes the JSON error message
// if ?debug is set it will response with the original errors message
func (d DefaultResp) Fail(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	resp := new(ErrResponse)

	w.Header().Add("Content-Type", "application/json")
	resp.Error.Msg = err.Error()
	if errors.Is(err, errors.ErrBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
	} else if errors.Is(err, errors.ErrUnauthorized) {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		file := errors.Caller()

		// hlog.FromRequest(r).Error().Err(err).Str("file", file).Send()

		w.WriteHeader(http.StatusInternalServerError)
		resp.Error.Codes = append(resp.Error.Codes, Error{
			Code: "internal_server_error",
		})

		if len(r.URL.Query()["debug"]) == 0 {
			resp.Error.File = file
		} else {
			resp.Error.Msg = http.StatusText(http.StatusInternalServerError)
		}
	}

	var ce *errors.CodeErr
	errs := err
	for {
		if errors.As(errs, &ce) {
			if ce.Code != "" {
				resp.Error.Codes = append(resp.Error.Codes, Error{
					Code: ce.Code,
					Msg:  ce.Msg,
					Args: ce.Args,
				})
			}
			errs = errors.Unwrap(ce)
			continue
		}

		break
	}

	d.JSON(w, r, resp)
}

// FailF same as Fail, but with error format
func (d DefaultResp) Failf(w http.ResponseWriter, r *http.Request, format string, a ...interface{}) {
	d.Fail(w, r, fmt.Errorf(format, a...))
}

// JSON writes the content of the param data as JSON.
// if ?pretty is present, it will pretty print the response.
func (d DefaultResp) JSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	if len(r.URL.Query()["pretty"]) != 0 {
		e.SetIndent(" ", " ")
	}
	if err := e.Encode(data); err != nil {
		d.Fail(w, r, fmt.Errorf("could not write json response; %w", err))
	}
}
