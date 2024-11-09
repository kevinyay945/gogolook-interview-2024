package di

import "gogolook/http"

type DI struct {
	HttpServer http.ServerInterface
}

func NewDI(httpServer http.ServerInterface) DI {
	return DI{HttpServer: httpServer}
}

func InitializeDI() DI {
	return NewDI(http.NewRestful())
}
