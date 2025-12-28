package main

import (
	maindb "iam-service/infrastructure/databases/maindb"
	"iam-service/interfaces/job_worker/worker"
)

func main() {
	maindb.InitConnection()

	worker.StartListen()
}
