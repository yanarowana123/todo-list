package inmemory

import (
	"context"
	"github.com/google/uuid"
	"log"
	"todo/internal/entity"
)

type Repository struct {
	m map[uuid.UUID]entity.Task
}

func NewRepository() Repository {
	m := make(map[uuid.UUID]entity.Task)
	return Repository{m}
}

func (r Repository) Create(ctx context.Context, t entity.Task) (uuid.UUID, error) {
	r.m[t.Id()] = t
	log.Println(r.m)
	return t.Id(), nil
}

func (r Repository) Update(ctx context.Context, id uuid.UUID) error {
	return nil
}
