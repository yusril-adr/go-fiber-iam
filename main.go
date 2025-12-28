package main

import (
	"github.com/robfig/cron/v3"

	"iam-service/infrastructure/databases/maindb"
	asyncqintegration "iam-service/infrastructure/integrations/asyncq"
	"iam-service/interfaces/http"
	authJobs "iam-service/modules/iam/auth/jobs"
)

func main() {
	maindb.InitConnection()

	asyncqintegration.InitClient()

	initScheduler()

	http.StartListen()
}

func initScheduler() {
	scheduler := cron.New()

	scheduler.AddFunc("0 0 * * *", authJobs.ClearExpiredToken)

	scheduler.Start()
}
