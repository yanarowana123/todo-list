package service

import (
	"context"
	"github.com/google/uuid"
	"time"
	"todo/internal/entity"
	"todo/internal/repository"
	"todo/internal/repository/converter"
)

type TaskRepository interface {
	Create(ctx context.Context, task repository.Task) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID) error
	//Delete(id uuid.UUID) error
	//DoTask(id uuid.UUID) error
	//List(status entity.Status) error
}

// TaskService Сервис является "фасадом" который взаимодействует с другими сервисами и вызывает кор-логику, реализованную в entity
type TaskService struct {
	r TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return TaskService{r}
}

func (t TaskService) Create(ctx context.Context, title string, activeAt time.Time) (uuid.UUID, error) {
	task, err := entity.NewTask(title, activeAt)
	if err != nil {
		return uuid.Nil, err
	}
	repoTask := converter.DomainTaskToRepo(*task)
	return t.r.Create(ctx, repoTask)
}

func (t TaskService) Update(ctx context.Context, id uuid.UUID) error {
	return t.r.Update(ctx, id)
}
