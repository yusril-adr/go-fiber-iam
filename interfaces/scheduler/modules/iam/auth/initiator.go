package auth

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	authRepository "iam-service/modules/iam/auth/repositories"
)

func RegisterScheduler(scheduler *cron.Cron) {
	scheduler.AddFunc("0 0 * * *", func() {
		logrus.Info("Deleting Tokens with expired time less than today...")
		today := time.Now()
		authRepository.DeleteTokenWithExpiredAt(today, nil)
	})
}
