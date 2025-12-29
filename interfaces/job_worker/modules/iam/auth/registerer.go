package auth

import (
	"github.com/hibiken/asynq"

	authhandler "iam-service/interfaces/job_worker/modules/iam/auth/handler"
	authJobs "iam-service/modules/iam/auth/jobs"
)

func RegisterHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(authJobs.TypeClearExpiredToken, authhandler.ClearExpiredToken)
}
