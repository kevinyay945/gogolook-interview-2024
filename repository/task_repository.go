package repository

import (
	"context"
	"gogolook/domain"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, task domain.Task) (domain.Task, error)
	UpdateByID(ctx context.Context, id string, task domain.Task) (domain.Task, error)
	DeleteByID(ctx context.Context, id string) error
}
