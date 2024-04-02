package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
	"unicode/utf8"
)

type Status int

const (
	active Status = iota
	done
)

// Task Главная сущность приложения
type Task struct {
	Id        uuid.UUID
	Title     string
	Status    Status
	ActiveAt  time.Time
	CreatedAt time.Time
}

func NewTask(title string, activeAt time.Time) (*Task, error) {
	// Учитываем мултьтибайтность
	if utf8.RuneCountInString(title) > 4 {
		return nil, errors.New("title length must not exceed 200 characters")
	}

	return &Task{uuid.New(), title, active, activeAt, time.Now()}, nil
}

func (t *Task) Done() error {
	t.Status = done
	return nil
}

//// Id геттеры для обеспечения имутабильности
//// приемник ссылочного типа потому что пункт 3 из https://youtu.be/G-lhh_1XNcI?list=LL&t=1540 (не знаю насколько можно доверять, но интересный доклад)
//func (t *Task) Id() uuid.UUID {
//	return t.id
//}
