package mongodb

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"todo/internal/repository"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return Repository{db}
}

func (r Repository) Create(ctx context.Context, t repository.Task) (uuid.UUID, error) {
	res, err := r.db.Collection("tasks").InsertOne(ctx, t)
	log.Println(res, err)
	return uuid.Nil, nil
}

func (r Repository) Update(ctx context.Context, id uuid.UUID) error {
	return nil
}
