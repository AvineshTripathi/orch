package task

import "github.com/oklog/ulid/v2"

type Task struct {
	ID    string
	URL   string
	Retry int
}

func NewTask(url string) *Task {
	return &Task{
		ID:    ulid.Make().String(),
		URL:   url,
		Retry: 0,
	}
}

func (t *Task) GetID() string {
	return t.ID
}

func (t *Task) GetURL() string {
	return t.URL
}

func (t *Task) GetRetry() int {
	return t.Retry
}

func (t *Task) IncrementRetry() {
	t.Retry++
}
