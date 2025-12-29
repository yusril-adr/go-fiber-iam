package main

import (
	"iam-service/infrastructure/databases/maindb"
	asyncqintegration "iam-service/infrastructure/integrations/asyncq"
	"iam-service/interfaces/http"
	"iam-service/interfaces/scheduler"
)

func main() {
	maindb.InitConnection()

	asyncqintegration.InitClient()

	scheduler.Init()

	http.StartListen()
}
