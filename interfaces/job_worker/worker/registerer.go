package worker

import (
	"github.com/hibiken/asynq"

	authWorker "iam-service/interfaces/job_worker/modules/iam/auth"
)

func RegisterHandlers(mux *asynq.ServeMux) {
	authWorker.RegisterHandlers(mux)
}
