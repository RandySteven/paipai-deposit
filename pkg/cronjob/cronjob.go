package cronjob_client

import (
	"context"
	"time"

	"github.com/RandySteven/paipai-deposit/configs"
	"github.com/robfig/cron/v3"
)

type (
	Scheduler interface {
		Run(ctx context.Context) error
		Stop(ctx context.Context) error
	}

	scheduler struct {
		cronJob *cron.Cron
	}
)

func NewScheduler(config *configs.Config) (Scheduler, error) {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")

	return &scheduler{
		cronJob: cron.New(cron.WithSeconds(), cron.WithLocation(jakartaTime)),
	}, nil
}
