package workers

import (
	"fmt"
	"orch/provisioner/task"
	"sync"
)

type SideLineWorker struct {
	ErrChan chan task.Task
	wg      *sync.WaitGroup
	quit    chan bool
}

func NewSideLineWorker(errChan chan task.Task, wg *sync.WaitGroup) *SideLineWorker {
	return &SideLineWorker{
		ErrChan: errChan,
		wg:      wg,
		quit:    make(chan bool),
	}
}

func (s *SideLineWorker) StartSideLineWorker() {
	go func() {
		for {
			select {
			case task := <-s.ErrChan:
				// should put into the queue of messages with retry ++
				fmt.Printf("SideLineWorker handling failed task %d (Retry #%d)\n", task.GetID(), task.GetRetry())
				err := s.processFailedTask(task)
				if err != nil {
					fmt.Printf("SideLineWorker failed to process task %d: %s\n", task.GetID(), err)
				}
			case <-s.quit:
				fmt.Printf("SideLineWorker stopping...\n")
				s.wg.Done()
				return
			}
		}
	}()
}

func (s *SideLineWorker) StopSideLineWorker() {
	s.quit <- true
}

// should put the task in the queue again
func (s *SideLineWorker) processFailedTask(task task.Task) error {
	task.IncrementRetry()
	if task.GetRetry() > 3 {
		return fmt.Errorf("max retries exceeded for task %d", task.GetID())
	}
	fmt.Printf("Retrying task %d (Retry #%d)\n", task.GetID(), task.GetRetry())
	return nil
}
