package dto

import (
	"encoding/json"
	"strings"
	"time"
)

type ErrorResponse struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}

type TaskResponse struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	ActiveAt Datetime `json:"activeAt"`
}

type Datetime struct {
	time.Time
}

// UnmarshalJSON для десериализации в Datetime
func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}

// MarshalJSON для сериализации в строку формата 2006-01-01
func (t *Datetime) MarshalJSON() ([]byte, error) {
	formatted := t.Time.Format("2006-01-02")
	return json.Marshal(formatted)
}

// UpsertTaskSwagger Для отображения в свагере, пока не нашел способ описать Datetime
type UpsertTaskSwagger struct {
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt" format:"date"`
}

type UpsertTaskRequest struct {
	Title    string   `json:"title"`
	ActiveAt Datetime `json:"activeAt"`
}

type CreateTaskResponse struct {
	Id string `json:"id"`
}
