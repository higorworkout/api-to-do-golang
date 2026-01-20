package database

import (
	"context"
	"time"

	"github.com/higorworkout/todo-api/internal/domain"
	"gorm.io/gorm"
)

type TaskModel struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Completed bool
	CreatedAt time.Time
}

type TaskGormRepository struct {
	db *gorm.DB
}

func NewTaskGormRepository(db *gorm.DB) *TaskGormRepository {
	db.AutoMigrate(&TaskModel{})
	return &TaskGormRepository{db: db}
}

func (r *TaskGormRepository) Create(ctx context.Context, task *domain.Task) error {
	model := TaskModel(*task)
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *TaskGormRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	var model TaskModel
	if err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	task := domain.Task(model)
	return &task, nil
}


func (r *TaskGormRepository) FindAll(ctx context.Context) ([]domain.Task, error) {
	var models []TaskModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0, len(models))
	for _, m := range models {
		tasks = append(tasks, domain.Task(m))
	}

	return tasks, nil
}

func (r *TaskGormRepository) Update(ctx context.Context, task *domain.Task) error {
	return r.db.WithContext(ctx).
		Model(&TaskModel{}).
		Where("id = ?", task.ID).
		Updates(TaskModel(*task)).Error
}

func (r *TaskGormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&TaskModel{}, "id = ?", id).Error
}