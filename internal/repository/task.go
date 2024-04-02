package repository

import (
	"time"
)

type Task struct {
	Id        string `bson:"_id"`
	Title     string
	Status    int
	ActiveAt  time.Time
	CreatedAt time.Time
}
