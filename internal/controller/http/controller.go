package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type TaskService interface {
	Create(ctx context.Context, title string, activeAt time.Time) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID) error
}

type Controller struct {
	s TaskService
}

func NewController(s TaskService) Controller {
	return Controller{s}
}

func (c Controller) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var createTaskRequest CreateTaskRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&createTaskRequest)
	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	id, err := c.s.Create(ctx, createTaskRequest.Title, createTaskRequest.ActiveAt.Time)

	if err != nil {
		c.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(CreateTaskResponse{id})
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request) {
	//return c.taskService.Update(id)
}

func (c Controller) respondWithError(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{statusCode, errorMsg})
}
