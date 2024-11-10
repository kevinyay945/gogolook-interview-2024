package domain

import (
	"github.com/google/uuid"
	"time"
)

type TaskStatus int

const (
	TASK_UNCOMPLETED TaskStatus = 0
	TASK_COMPLETED   TaskStatus = 1
)

type Task struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Status    TaskStatus
}

func NewTask(id uuid.UUID, name string, status TaskStatus) Task {
	return Task{ID: id, Name: name, Status: status}
}
