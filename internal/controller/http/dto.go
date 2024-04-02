package http

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

//Здесь описаны сущности слоя транспорта. В идеале каждый слой должен оперировать своими сущностями

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}

type CreateTaskRequest struct {
	Title    string
	ActiveAt Datetime
}

type CreateTaskResponse struct {
	Id uuid.UUID
}

type ErrorResponse struct {
	ErrorCode int
	Message   string
}
