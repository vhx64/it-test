package ports

import (
	"it-test/app"
	"it-test/pkg/server/httperr"
)

type Validator interface {
	Valid() *httperr.ErrorMessageBody
}

var (
	_ ServerInterface = HTTPServer{}
)

type HTTPServer struct {
	app *app.Application
}

func NewHTTPServer(a *app.Application) HTTPServer {
	return HTTPServer{app: a}
}
