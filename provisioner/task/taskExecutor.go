package task

import (
	"fmt"
	"os"
	"plugin"
	"sync"

	"gopkg.in/yaml.v2"
)

type TaskExecutor struct {
	mu    sync.Mutex
	plugins map[string]*plugin.Plugin
}

func NewTaskExecutor() *TaskExecutor {
	return &TaskExecutor{
		plugins: make(map[string]*plugin.Plugin),
	}
}

func (t *TaskExecutor) LoadPlugin(path string) error {

	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %v", err)
	}

	var pluginConfig map[string]string
	err = yaml.Unmarshal(file, &pluginConfig)
	if err != nil {
		return fmt.Errorf("failed to parse YAML file: %v", err)
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	for functionName, pluginPath := range pluginConfig {
		p, err := plugin.Open(pluginPath)
		if err != nil {
			return fmt.Errorf("failed to load plugin '%s' from path '%s': %v", functionName, pluginPath, err)
		}
		t.plugins[functionName] = p
	}
	return nil
}

func (t *TaskExecutor) Execute(task *Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	  
	for _, p := range t.plugins {

		runFunc, err := p.Lookup(task.Plugin)
		if err != nil {
			return err
		}

		// donot need for now
		// config := models.Config{
		// 	Cfg: task.Config,
		// 	LogFile:  "output.log",
		// 	LogLevel: 2,
		// }

		fn, ok := runFunc.(func(args ...interface{}) error)
		if !ok {
			return fmt.Errorf("unexpected type from module symbol")
		}

		err = fn(task.Config)
		if err != nil {
			return fmt.Errorf("error while running the plugin: %v", err)
		}
	}

	return nil
}
