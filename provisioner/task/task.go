package task

import (
	"github.com/oklog/ulid/v2"
)

type Task struct {
	Id     string
	Plugin string
	Data   string
	Config any
	Retry  int
}

func NewTask(plugin, data string, cfg any) *Task {
	return &Task{
		Id:    ulid.Make().String(),
		Plugin:   plugin,
		Data:  data,
		Config: cfg,
		Retry: 0,
	}
}

func (t *Task) GetID() string {
	return t.Id
}

func (t *Task) GetPluginName() string {
	return t.Plugin
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
