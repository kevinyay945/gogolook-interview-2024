package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gogolook/domain"
	"gogolook/lib/pg"
	"gorm.io/gorm"
	"time"
)

type PostgresqlTaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &PostgresqlTaskRepository{db: db}
}

func (t *PostgresqlTaskRepository) FindAll(ctx context.Context) ([]domain.Task, error) {
	var tasks []pg.Task
	result := t.db.WithContext(ctx).Model(&pg.Task{}).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	var taskList []domain.Task
	for _, task := range tasks {
		parse, err := uuid.Parse(task.ID)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, domain.Task{
			ID:        parse,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			Name:      task.Name,
			Status:    domain.TaskStatus(task.Status),
		})
	}
	return taskList, nil
}

func (t *PostgresqlTaskRepository) Create(ctx context.Context, task domain.Task) (domain.Task, error) {

	pgTask := pg.Task{
		ID:        task.ID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      task.Name,
		Status:    int(task.Status),
	}
	result := t.db.WithContext(ctx).Create(&pgTask)
	if result.Error != nil {
		return domain.Task{}, result.Error
	}
	if result.RowsAffected != 1 {
		return domain.Task{}, fmt.Errorf("create: did not create task")
	}

	parseUUID, err := uuid.Parse(pgTask.ID)
	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{
		ID:        parseUUID,
		CreatedAt: pgTask.CreatedAt,
		UpdatedAt: pgTask.UpdatedAt,
		Name:      pgTask.Name,
		Status:    domain.TaskStatus(pgTask.Status),
	}, nil
}

func (t *PostgresqlTaskRepository) UpdateByID(ctx context.Context, id string, task domain.Task) (domain.Task, error) {
	result := t.db.WithContext(ctx).Model(&pg.Task{}).
		Where("id = ?", id).
		Updates(pg.Task{
			UpdatedAt: time.Now(),
			Name:      task.Name,
			Status:    int(task.Status),
		})
	if result.Error != nil {
		return domain.Task{}, result.Error
	}
	if result.RowsAffected != 1 {
		return domain.Task{}, fmt.Errorf("updateByID: did not update task")
	}
	return task, nil
}

func (t *PostgresqlTaskRepository) DeleteByID(ctx context.Context, id string) error {
	result := t.db.WithContext(ctx).Model(&pg.Task{}).Where("id = ?", id).
		Delete(&pg.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("deleteByID: did not delete task")
	}
	return nil
}
