package repository

import (
	"context"
	"time"
)

type Task struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Status    int
}

type TaskRepository interface {
	FindAll(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, task Task) (Task, error)
	UpdateByID(ctx context.Context, id string, task Task) (Task, error)
	DeleteByID(ctx context.Context, id string) error
}
