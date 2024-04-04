package service

import (
	"context"
	"errors"
	"strings"
	"time"
	"todo/internal/entity"
	"todo/internal/service/dto"
)

type TaskRepository interface {
	Create(ctx context.Context, task entity.Task) (string, error)
	Update(ctx context.Context, task entity.Task) error
	Find(ctx context.Context, id string) *entity.Task
	Delete(ctx context.Context, task entity.Task) error
	List(ctx context.Context, status entity.Status) ([]entity.Task, error)
}

// TaskService Сервис является "фасадом" который взаимодействует с другими сервисами и вызывает кор-логику, реализованную в entity
type TaskService struct {
	r TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return TaskService{r}
}

func (s TaskService) Create(ctx context.Context, createTaskRequest dto.UpsertTask) (string, error) {
	task, err := entity.NewTask(createTaskRequest.Title, createTaskRequest.ActiveAt)
	if err != nil {
		return "", err
	}

	return s.r.Create(ctx, *task)
}

func (s TaskService) Update(ctx context.Context, id string, updateTaskRequest dto.UpsertTask) error {
	task := s.r.Find(ctx, id)

	if task == nil {
		return errors.New("task not found")
	}

	err := task.Update(updateTaskRequest.Title, updateTaskRequest.ActiveAt)

	if err != nil {
		return err
	}

	return s.r.Update(ctx, *task)
}

func (s TaskService) Delete(ctx context.Context, id string) error {
	task := s.r.Find(ctx, id)

	if task == nil {
		return errors.New("task not found")
	}

	return s.r.Delete(ctx, *task)
}

func (s TaskService) Done(ctx context.Context, id string) error {
	task := s.r.Find(ctx, id)

	if task == nil {
		return errors.New("task not found")
	}

	err := task.Done()

	if err != nil {
		return err
	}

	return s.r.Update(ctx, *task)
}

func (s TaskService) List(ctx context.Context, status entity.Status) ([]dto.Task, error) {
	var list []dto.Task
	tasks, err := s.r.List(ctx, status)
	if err != nil {
		return list, err
	}
	for _, task := range tasks {
		var title string
		if isWeekend(task.ActiveAt) {
			var sb strings.Builder
			sb.WriteString("ВЫХОДНОЙ - ")
			sb.WriteString(task.Title)
			title = sb.String()
		} else {
			title = task.Title
		}

		list = append(list, dto.Task{task.Id, title, task.ActiveAt})
	}

	return list, nil
}

func isWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
