package jobs

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/jobs/tasks"
	"go.uber.org/zap"
)

type asyncClient struct {
	*asynq.Client
	*zap.Logger
}

func newAsyncClient(conf conf.Configuration, l *zap.Logger) (asyncClient, error) {
	cl := asyncClient{
		asynq.NewClient(asynq.RedisClientOpt{Addr: conf.Jobs.RedisAddress}), l,
	}

	tasks := cl.addTasks()
	for _, task := range tasks {
		info, err := cl.Enqueue(task)
		if err != nil {
			return asyncClient{}, err
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}
	return cl, nil
}

func (ac *asyncClient) addTasks() []*asynq.Task {
	t := []*asynq.Task{
		tasks.NewImageResizeTask(),
	}
	return t
}
