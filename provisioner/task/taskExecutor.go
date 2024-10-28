// Package task demonstrates how to create a modular Go program using plugins.
// It defines the Task interface and TaskExecutor struct, providing methods for loading and executing plugins.
package task

import (
	"fmt"
	"plugin"
	"sync"

	"github.com/AvineshTripathi/orch/provisioner/models"
)

type TaskExecutor struct {
	mu    sync.Mutex
	tasks []*plugin.Plugin
}

func NewTaskExecutor() *TaskExecutor {
	return &TaskExecutor{}
}

func (t *TaskExecutor) LoadPlugin(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}
	t.tasks = append(t.tasks, p)
	return nil
}

func (t *TaskExecutor) Execute(task *Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, p := range t.tasks {

		runFunc, err := p.Lookup("Run")
		if err != nil {
			return err
		}

		config := models.Config{
			Cfg: task.Config,
			LogFile:  "output.log",
			LogLevel: 2,
		}

		fn, ok := runFunc.(func(cfg models.Config) error)
		if !ok {
			return fmt.Errorf("unexpected type from module symbol")
		}

		err = fn(config)
		if err != nil {
			fmt.Errorf("Error while running the plugin: %v", err)
		}
	}

	return nil
}
