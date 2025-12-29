package scheduler

import (
	"github.com/robfig/cron/v3"

	"iam-service/interfaces/scheduler/modules/iam/auth"
)

func Init() {
	scheduler := cron.New()

	auth.RegisterScheduler(scheduler)

	scheduler.Start()
}
