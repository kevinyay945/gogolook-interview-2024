package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"gogolook/domain"
	"gogolook/service"
	"net/http"
)

type RestfulServer struct {
	taskService service.TaskService
}

func NewRestfulServer(taskService service.TaskService) ServerInterface {
	return &RestfulServer{taskService: taskService}
}

func (r *RestfulServer) GetTasks(ctx echo.Context) error {
	tasks, err := r.taskService.FindAllTasks(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var output []Task
	for _, task := range tasks {
		output = append(output, Task{
			Id:        task.ID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			Name:      task.Name,
			Status:    TaskStatus(task.Status),
		})
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func (r *RestfulServer) PostTask(ctx echo.Context) error {
	input := struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}{}
	id, err := uuid.NewRandom()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = ctx.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createTask, err := r.taskService.CreateTask(ctx.Request().Context(),
		id, input.Name, domain.TaskStatus(input.Status),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task := Task{
		Id:        createTask.ID,
		CreatedAt: createTask.CreatedAt,
		UpdatedAt: createTask.UpdatedAt,
		Name:      createTask.Name,
		Status:    TaskStatus(createTask.Status),
	}
	return ctx.JSON(http.StatusCreated, task)
}

func (r *RestfulServer) DeleteTask(ctx echo.Context, id openapitypes.UUID) error {
	err := r.taskService.DeleteTaskByID(ctx.Request().Context(), id.String())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (r *RestfulServer) PutTask(ctx echo.Context, id openapitypes.UUID) error {
	task := Task{}
	err := ctx.Bind(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedTask, err := r.taskService.UpdateTaskByID(ctx.Request().Context(), id.String(),
		domain.NewTask(id, task.Name, domain.TaskStatus(task.Status)),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, updatedTask)
}
