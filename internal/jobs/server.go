package jobs

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/hugo.rojas/custom-api/internal/jobs/tasks"
)

func newAsyncServer(conf conf.Configuration, service iface.Service) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: conf.Jobs.RedisAddress},
		asynq.Config{
			Concurrency: conf.Jobs.Concurrency,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor(service))

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run async server %v", err)
	}
}
