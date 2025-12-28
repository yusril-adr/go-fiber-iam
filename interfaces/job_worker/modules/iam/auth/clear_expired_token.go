package auth

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	authService "iam-service/modules/iam/auth/services"
)

func ClearExpiredTokenHandler(ctx context.Context, t *asynq.Task) error {
	logrus.Info("Clearing expired token...")
	authService.ClearExpiredToken()

	return nil
}
