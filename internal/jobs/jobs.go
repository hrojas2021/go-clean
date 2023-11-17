package jobs

import (
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"go.uber.org/zap"
)

func InitJobsQueue(conf conf.Configuration, l *zap.Logger, service iface.Service) {
	if conf.Jobs.Enabled {
		client, err := newAsyncClient(conf, l)
		if err != nil {
			l.Error("could not create the async client", zap.Error(err))
		}
		defer client.Close()
		go newAsyncServer(conf, service)
	}
}
