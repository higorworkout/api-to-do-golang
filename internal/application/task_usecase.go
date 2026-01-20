package application

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/higorworkout/todo-api/internal/domain"
)

type TaskUseCase struct {
	repo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}


func (uc *TaskUseCase) CreateTask(ctx context.Context, title string) (*domain.Task, error) {
	task := &domain.Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: salvar no banco
	go func() {
		defer wg.Done()
		_ = uc.repo.Create(ctx, task)
	}()

	// Goroutine 2: simular evento / log / WoW no futuro
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
	}()

	wg.Wait()
	return task, nil
}

func (uc *TaskUseCase) ListTasks(ctx context.Context) ([]domain.Task, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, task *domain.Task) error {
	return uc.repo.Update(ctx, task)
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}