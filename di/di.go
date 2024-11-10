package di

import (
	"gogolook/http"
	"gogolook/lib/pg"
	"gogolook/repository"
	"gogolook/service"
)

type DI struct {
	HttpServer http.ServerInterface
}

func NewDI(httpServer http.ServerInterface) DI {
	return DI{HttpServer: httpServer}
}

func InitializeDI() DI {
	db := pg.GetDB()
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	server := http.NewRestfulServer(taskService)
	return NewDI(server)
}
