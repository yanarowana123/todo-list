package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateTask struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string
	Status    int
	ActiveAt  time.Time `bson:"activeAt"`
	CreatedAt time.Time `bson:"createdAt"`
}

type Task struct {
	Id        string `bson:"_id"`
	Title     string
	Status    int
	ActiveAt  time.Time
	CreatedAt time.Time
}
