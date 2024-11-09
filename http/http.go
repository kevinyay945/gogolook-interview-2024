package http

import (
	"github.com/labstack/echo/v4"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

type RestfulServer struct{}

func NewRestful() ServerInterface {
	return &RestfulServer{}
}

func (r RestfulServer) GetTasks(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r RestfulServer) PostTask(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r RestfulServer) DeleteTask(ctx echo.Context, id openapitypes.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r RestfulServer) PutTask(ctx echo.Context, id openapitypes.UUID) error {
	//TODO implement me
	panic("implement me")
}
