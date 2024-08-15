package workers

import (
	"fmt"
	"orch/provisioner/task"
	"sync"
)

type Worker struct {
	ID        int
	TaskChan  chan task.Task
	ErrChan   chan task.Task
	wg        *sync.WaitGroup
	quit      chan bool
}

func NewWorker(id int, taskChan chan task.Task, errChan chan task.Task, wg *sync.WaitGroup,) *Worker {
	return &Worker{
		ID:        id,
		TaskChan:  taskChan,
		ErrChan:   errChan,
		wg:        wg,
		quit:      make(chan bool),
	}
}

func (w *Worker) StartWorker() {
	go func() {
		for {
			select {
			case task := <- w.TaskChan:
				fmt.Printf("Worker %d processing task %d\n", w.ID, task.GetID())
			case <- w.quit:
				fmt.Printf("Worker %d stopping...\n", w.ID)
				w.wg.Done()
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.quit <- true
}

