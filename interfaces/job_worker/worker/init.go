package worker

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	asyncqintegration "iam-service/infrastructure/integrations/asyncq"
)

func StartListen() {
	asynqConfig, asyncqRedisConfig := asyncqintegration.GetConfig()

	srv := asynq.NewServer(
		asyncqRedisConfig,
		asynqConfig,
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	RegisterHandlers(mux)

	logrus.Info("Scheduler Worker server running ...")

	if err := srv.Run(mux); err != nil {
		panic(fmt.Sprintf("could not run server: %v", err))
	}
}
