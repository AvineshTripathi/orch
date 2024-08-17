package queue

import (
	"context"
	"encoding/json"
	"orch/provisioner/task"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Db        *redis.Client
	Ctx       context.Context
	QueueName string
}

func NewConnection() *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Client{
		Db:        client,
		Ctx:       context.Background(),
		QueueName: "random",
	}
}

func (client *Client) AddNewTask(task *task.Task) (int64, error) {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return -1, err
	}

	res := client.Db.ZAdd(client.Ctx, client.QueueName, redis.Z{
		Score: float64(task.Retry),
		Member: taskJson,
	})
	return res.Result()
}


func (client *Client) GetTasks() ([]string, error) {
	res := client.Db.ZRange(client.Ctx, client.QueueName, 0, -1)
	return res.Result()	
}