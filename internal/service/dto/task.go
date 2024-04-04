package dto

import (
	"time"
)

type UpsertTask struct {
	Title    string
	ActiveAt time.Time
}

type Task struct {
	Id       string
	Title    string
	ActiveAt time.Time
}
