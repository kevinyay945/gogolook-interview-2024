package pg

import "time"

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
