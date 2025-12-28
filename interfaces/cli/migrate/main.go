package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"iam-service/constants"
	"iam-service/infrastructure/databases/maindb"
)

func main() {
	// Command: up or down
	if len(os.Args) < 2 {
		log.Fatal("missing command: up | down")
	}

	command := os.Args[1]

	// Parse flags
	defaultDB := constants.DB_LIST_MAIN // default database
	dbName := flag.String("db", defaultDB, "database name to migrate")

	flag.CommandLine.Parse(os.Args[2:])

	if *dbName == "" {
		*dbName = defaultDB
	}

	switch command {
	case "up":
		fmt.Println("Migrating UP database:", *dbName)
		runUp(*dbName)

	case "down":
		fmt.Println("Migrating DOWN database:", *dbName)
		runDown(*dbName)

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("force requires a version number")
		}

		rawVersion := os.Args[2]
		version, err := strconv.Atoi(rawVersion)
		if err != nil {
			log.Fatal("invalid version number:", rawVersion)
		}

		fmt.Println("Migrating FORCE database:", *dbName)
		runForce(*dbName, version)

	default:
		log.Fatal("unknown command:", command)
	}

	fmt.Println("Migration completed.")
}

func runUp(db string) {
	switch db {
	case constants.DB_LIST_MAIN:
		maindb.MigrationUp()

	default:
		log.Fatal("unknown database:", db)
	}
}

func runDown(db string) {
	switch db {
	case constants.DB_LIST_MAIN:
		maindb.MigrationDown()

	default:
		log.Fatal("unknown database:", db)
	}
}

func runForce(db string, version int) {
	switch db {
	case constants.DB_LIST_MAIN:
		maindb.MigrationForce(version)

	default:
		log.Fatal("unknown database:", db)
	}
}
