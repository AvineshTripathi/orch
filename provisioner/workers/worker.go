package workers

import (
	"fmt"
	"log"
	"github.com/AvineshTripathi/orch/provisioner/task"
	"sync"
)

type Worker struct {
	ID       int
	TaskChan chan task.Task
	Executor *task.TaskExecutor
	ErrChan  chan task.Task
	wg       *sync.WaitGroup
	quit     chan bool
}

func NewWorker(id int, taskChan chan task.Task, errChan chan task.Task, taskExecutor *task.TaskExecutor, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:       id,
		TaskChan: taskChan,
		ErrChan:  errChan,
		wg:       wg,
		Executor: taskExecutor,
		quit:     make(chan bool),
	}
}

func (w *Worker) StartWorker() {
	go func() {
		for {
			select {
			case task := <-w.TaskChan:
				fmt.Printf("Worker %d processing task id %s, plugin %s", w.ID, task.GetID(), task.GetPluginName())
				err := w.Executor.Execute(&task)
				if err != nil {
					log.Println("Couldnot execute", err)
					w.ErrChan <- task
				}
			case <-w.quit:
				fmt.Printf("Worker %d stopping...\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.quit <- true
}
