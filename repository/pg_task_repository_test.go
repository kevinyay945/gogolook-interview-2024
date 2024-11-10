package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gogolook/domain"
	"gogolook/lib/pg"
	"gorm.io/gorm"
	"testing"
)

type TaskSuite struct {
	suite.Suite
	db       *gorm.DB
	taskRepo TaskRepository
}

// TestSuiteInitTask is only for development and useCase, remove it in production
func TestSuiteInitTask(t *testing.T) {
	t.Skip()
	suite.Run(t, new(TaskSuite))
}
func (t *TaskSuite) SetupTest() {
	t.db = pg.GetDBByConnectingString("postgresql://myuser:mypassword@127.0.0.1/mydb?sslmode=disable")
	t.taskRepo = NewTaskRepository(t.db)
}

func (t *TaskSuite) TearDownTest() {
}

func (t *TaskSuite) Test_create_task() {
	task := domain.NewTask(uuid.New(), "test", 1)
	create, err := t.taskRepo.Create(context.Background(), task)
	t.NoError(err)
	t.NotEmpty(create.ID)
}

func (t *TaskSuite) Test_get_all_task() {
	tasks, err := t.taskRepo.FindAll(context.Background())
	t.NoError(err)
	t.NotEmpty(tasks)
}

func (t *TaskSuite) Test_update_task_by_id() {
	task, err := t.taskRepo.UpdateByID(context.Background(), "02582465-340a-4269-876e-c49eb25acc9a", domain.Task{Name: "test", Status: 0})
	t.NoError(err)
	t.Equal(task.Status, 0)
}

func (t *TaskSuite) Test_delete_task_by_id() {
	err := t.taskRepo.DeleteByID(context.Background(), "02582465-340a-4269-876e-c49eb25acc9a")
	t.NoError(err)
}
