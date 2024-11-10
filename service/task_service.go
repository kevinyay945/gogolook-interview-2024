package service

import (
	"context"
	"github.com/google/uuid"
	"gogolook/domain"
	"gogolook/repository"
)

// TaskService provides task-related operations.
type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return TaskService{repo: repo}
}

func (s *TaskService) FindAllTasks(ctx context.Context) ([]domain.Task, error) {
	return s.repo.FindAll(ctx)
}

func (s *TaskService) CreateTask(ctx context.Context, id uuid.UUID, name string, status domain.TaskStatus) (domain.Task, error) {
	task := domain.NewTask(id, name, status)

	newTask, err := s.repo.Create(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return newTask, nil
}

func (s *TaskService) UpdateTaskByID(ctx context.Context, id string, task domain.Task) (domain.Task, error) {
	return s.repo.UpdateByID(ctx, id, task)
}

func (s *TaskService) DeleteTaskByID(ctx context.Context, id string) error {
	return s.repo.DeleteByID(ctx, id)
}
