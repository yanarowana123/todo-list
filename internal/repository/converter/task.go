package converter

import (
	"github.com/google/uuid"
	"todo/internal/entity"
	"todo/internal/repository"
)

// Этот пакет используется для преобразования сущностей из одного слоя в сущность другого слоя (внешние слои знают о внутренних)

func RepoTaskToDomain(t repository.Task) entity.Task {
	return entity.Task{uuid.Nil, t.Title, entity.Status(t.Status), t.ActiveAt, t.CreatedAt}
}

func DomainTaskToRepo(t entity.Task) repository.Task {
	return repository.Task{"f", t.Title, int(t.Status), t.ActiveAt, t.CreatedAt}
}
