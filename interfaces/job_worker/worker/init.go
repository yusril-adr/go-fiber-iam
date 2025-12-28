package worker

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	asyncqintegration "iam-service/infrastructure/integrations/asyncq"
	authWorker "iam-service/interfaces/job_worker/modules/iam/auth"
	authJobs "iam-service/modules/iam/auth/jobs"
)

func StartListen() {
	asynqConfig, asyncqRedisConfig := asyncqintegration.GetConfig()

	srv := asynq.NewServer(
		asyncqRedisConfig,
		asynqConfig,
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(authJobs.TypeClearExpiredToken, authWorker.ClearExpiredTokenHandler)

	logrus.Info("Scheduler Worker server running ...")

	if err := srv.Run(mux); err != nil {
		panic(fmt.Sprintf("could not run server: %v", err))
	}
}
