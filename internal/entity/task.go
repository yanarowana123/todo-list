package entity

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Status int

const (
	BrandNew Status = iota
	Done
)

// Task Главная сущность приложения
type Task struct {
	Id        string
	Title     string
	Status    Status
	ActiveAt  time.Time
	CreatedAt time.Time
}

func NewTask(title string, activeAt time.Time) (*Task, error) {
	err := validateTitle(title)
	if err != nil {
		return nil, errors.New("title length must not exceed 200 characters")
	}

	return &Task{"", title, BrandNew, activeAt, time.Now()}, nil
}

func (t *Task) Update(title string, activeAt time.Time) error {
	err := validateTitle(title)
	if err != nil {
		return errors.New("title length must not exceed 200 characters")
	}

	t.Title = title
	t.ActiveAt = activeAt
	return nil
}

func (t *Task) Done() error {
	t.Status = Done
	return nil
}

func validateTitle(title string) error {
	if utf8.RuneCountInString(title) > 200 {
		return errors.New("title length must not exceed 200 characters")
	}

	return nil
}
