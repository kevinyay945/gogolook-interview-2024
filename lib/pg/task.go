package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gogolook/repository"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        string    `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`
}

func (Task) TableName() string {
	return "tasks"
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &TaskRepository{db: db}
}

func (t TaskRepository) FindAll(ctx context.Context) ([]repository.Task, error) {
	var tasks []Task
	result := t.db.WithContext(ctx).Model(&Task{}).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	var taskList []repository.Task
	for _, task := range tasks {
		taskList = append(taskList, repository.Task{
			ID:        task.ID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			Name:      task.Name,
			Status:    task.Status,
		})
	}
	return taskList, nil
}

func (t TaskRepository) Create(ctx context.Context, task repository.Task) (repository.Task, error) {
	uuidString := task.ID
	if uuidString == "" {
		uuidV4, err := uuid.NewRandom()
		if err != nil {
			return repository.Task{}, err
		}
		uuidString = uuidV4.String()
	}

	pgTask := Task{
		ID:        uuidString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      task.Name,
		Status:    task.Status,
	}
	result := t.db.WithContext(ctx).Create(&pgTask)
	if result.Error != nil {
		return repository.Task{}, result.Error
	}
	if result.RowsAffected != 1 {
		return repository.Task{}, fmt.Errorf("create: did not create task")
	}
	return repository.Task{
		ID:        pgTask.ID,
		CreatedAt: pgTask.CreatedAt,
		UpdatedAt: pgTask.UpdatedAt,
		Name:      pgTask.Name,
		Status:    pgTask.Status,
	}, nil
}

func (t TaskRepository) UpdateByID(ctx context.Context, id string, task repository.Task) (repository.Task, error) {
	result := t.db.WithContext(ctx).Model(&Task{}).
		Where("id = ?", id).
		Updates(Task{
			UpdatedAt: time.Now(),
			Name:      task.Name,
			Status:    task.Status,
		})
	if result.Error != nil {
		return repository.Task{}, result.Error
	}
	if result.RowsAffected != 1 {
		return repository.Task{}, fmt.Errorf("updateByID: did not update task")
	}
	return task, nil
}

func (t TaskRepository) DeleteByID(ctx context.Context, id string) error {
	result := t.db.WithContext(ctx).Model(&Task{}).Where("id = ?", id).Delete(&Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("deleteByID: did not delete task")
	}
	return nil
}
