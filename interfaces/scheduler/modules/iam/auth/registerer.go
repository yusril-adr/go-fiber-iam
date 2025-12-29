package auth

import (
	"github.com/robfig/cron/v3"

	authJob "iam-service/modules/iam/auth/jobs"
)

func RegisterScheduler(scheduler *cron.Cron) {
	scheduler.AddFunc("0 0 * * *", authJob.ClearExpiredToken)
	scheduler.AddFunc("* * * * *", authJob.ClearExpiredToken)
}
