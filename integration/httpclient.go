package integration

import (
	"net/http"
	"time"

	"github.com/federicoleon/go-httpclient/gohttp"
	"github.com/federicoleon/go-httpclient/gomime"
)

func getHTTPClient(jwt string) gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	headers.Set("Authorization", "Bearer "+jwt)

	return gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(4 * time.Second).
		SetResponseTimeout(6 * time.Second).
		Build()
}
