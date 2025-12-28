package jobs

import (
	asyncqintegration "iam-service/infrastructure/integrations/asyncq"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var (
	TypeClearExpiredToken = "iam:auth:clear_expired_token"
)

func ClearExpiredToken() {
	task := asynq.NewTask(TypeClearExpiredToken, nil, asynq.MaxRetry(10))

	logrus.Info("Enqueued ", TypeClearExpiredToken, "...")
	asyncqintegration.Client.Enqueue(task)
}
