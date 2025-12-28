air-init:
	go tool air init

air:
	go tool air

dev:
	go run main.go

job-worker:
	go run interfaces/job_worker/main.go

migrate-up:
	go run interfaces/cli/migrate/main.go up --db=${db}

migrate-down:
	go run interfaces/cli/migrate/main.go down --db=${db}

migrate-force:
	go run interfaces/cli/migrate/main.go force "${version}" --db=${db}

seed:
	go run interfaces/cli/seeder/main.go ${module} --db=${db}
