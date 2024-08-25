package task

import (
	"github.com/oklog/ulid/v2"
)

type Task struct {
	Id    string
	Url   string
	Data  string
	Retry int
}

func NewTask(url, data string) *Task {
	return &Task{
		Id:    ulid.Make().String(),
		Url:   url,
		Data:  data,
		Retry: 0,
	}
}

func (t *Task) GetID() string {
	return t.Id
}

func (t *Task) GetURL() string {
	return t.Url
}

func (t *Task) GetRetry() int {
	return t.Retry
}

func (t *Task) GetData() string {
	return t.Data
}

func (t *Task) IncrementRetry() {
	t.Retry++
}
