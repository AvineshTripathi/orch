package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AvineshTripathi/orch/config"
	"github.com/AvineshTripathi/orch/provisioner/task"

	"github.com/redis/go-redis/v9"
)

const (
	LIMIT        = int64(10)
	POLLING_TIME = 2 * time.Second
)

type Client struct {
	Db        *redis.Client
	Ctx       context.Context
	TaskChan  chan task.Task
	ErrChan   chan task.Task
	QueueName string
	quit      chan bool
}

func NewConnection() *Client {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
		DB:       0,
	})

	// Create a new client object
	c := &Client{
		Db:        client,
		Ctx:       context.Background(),
		QueueName: config.RedisQueue,
		quit:      make(chan bool),
	}
	return c
}

func (client *Client) ConfigureTaskChannel(taskChan, errChan chan task.Task) {
	client.TaskChan = taskChan
	client.ErrChan = errChan
}

func (client *Client) AddNewTask(task *task.Task) (int64, error) {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return -1, err
	}

	res := client.Db.ZAdd(client.Ctx, config.RedisQueue, redis.Z{
		Score:  float64(task.GetRetry()),
		Member: taskJson,
	})
	return res.Result()
}

func (client *Client) DeleteTask(task *task.Task) (int64, error) {
	// taskJson, err := json.Marshal(task)
	// if err != nil {
	// 	return -1, err
	// }

	// dara = redis.Z{
	// 	Score:  float64(task.GetRetry()),
	// 	Member: taskJson,
	// }

	res := client.Db.ZRem(client.Ctx, client.QueueName, task)

	return res.Result()
}

func (client *Client) GetTasksWithPagination(offset, limit int64) ([]string, error) {

	res, err := client.Db.ZRange(client.Ctx, client.QueueName, offset, offset+limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("error fetching task: %v", err)
	}

	if len(res) > 0 {
		_, err = client.Db.ZRem(client.Ctx, client.QueueName, convertToInterfaceSlice(res)...).Result()
		if err != nil {
			return nil, fmt.Errorf("error deleting task")
		}
	}

	return res, err
}

func convertToInterfaceSlice(strings []string) []interface{} {
	interfaces := make([]interface{}, len(strings))
	for i, v := range strings {
		interfaces[i] = v
	}
	return interfaces
}

func (client *Client) ProcessErrorConitnuously() {

	go func() {
		for {
			select {
			case <-client.quit:
				fmt.Println("Stopping task fetching...")
				return
			default:
				for task := range client.ErrChan {
					_, err := client.AddNewTask(&task)
					if err != nil {
						fmt.Println("Error queuing failed task...")
						continue
					}
				}
			}
		}
	}()
}

func (client *Client) ProcessTasksContinuously() {

	var tsk task.Task
	go func() {
		var offset int64
		for {
			select {
			case <-client.quit:
				fmt.Println("Stopping task fetching...")
				return
			default:
				tasks, err := client.GetTasksWithPagination(offset, LIMIT)
				if err != nil {
					fmt.Printf("Error : %v\n", err)
					time.Sleep(POLLING_TIME)
					continue
				}

				if len(tasks) == 0 {
					offset = 0
					time.Sleep(POLLING_TIME)
					continue
				}

				for _, taskJson := range tasks {
					if taskJson == "" {
						continue
					}
					err := json.Unmarshal([]byte(taskJson), &tsk)
					if err != nil {
						fmt.Printf("Error Unmarshalling tasks: %v\n", err)
						time.Sleep(POLLING_TIME)
						continue
					}
					client.TaskChan <- tsk

					// remove key from queue storage
					// _, err = client.DeleteTask(&task)
					// if err != nil {
					// 	fmt.Println("Error deleting task...")
					// 	continue
					// }
				}

				offset += LIMIT
			}

		}
	}()
}

func (client *Client) StopClient() {
	client.quit <- true
}
