package maindb

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"iam-service/infrastructure/config"
	"iam-service/infrastructure/databases/helpers/postgresql"
)

// Exported for Database Connection
var (
	Connection *gorm.DB
	DBProvider string = "postgres"
	DBParams   postgresql.DBParams
	DBUrl      string
)

func InitConnection() {
	maxIdleConn, errIdle := strconv.Atoi(config.MAIN_DB_MAX_IDLE_CONNS)
	if errIdle != nil {
		panic(errIdle.Error())
	}

	maxOpenConn, errOpen := strconv.Atoi(config.MAIN_DB_MAX_OPEN_CONNS)
	if errOpen != nil {
		panic(errOpen.Error())
	}

	DBParams = postgresql.DBParams{
		Host:        config.MAIN_DB_HOST,
		Port:        config.MAIN_DB_PORT,
		Username:    config.MAIN_DB_USER,
		Password:    config.MAIN_DB_PASS,
		Name:        config.MAIN_DB_NAME,
		SSLMode:     config.MAIN_DB_SSL_MODE,
		TimeZome:    config.MAIN_DB_TIMEZONE,
		MaxIdleConn: maxIdleConn,
		MaxOpenConn: maxOpenConn,
	}

	Connection = postgresql.InitDbConnection(DBParams)
	DBUrl = postgresql.GetUrl(DBParams)
}

func initMigration() *migrate.Migrate {
	rawDb, err := Connection.DB()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(rawDb, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/main_db",
		DBProvider,
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func MigrationUp() {
	InitConnection()
	m := initMigration()

	if err := m.Up(); err != nil {
		panic(err)
	}
}

func MigrationDown() {
	InitConnection()
	m := initMigration()

	if err := m.Down(); err != nil {
		panic(err)
	}
}

func MigrationForce(version int) {
	InitConnection()
	m := initMigration()

	if err := m.Force(version); err != nil {
		panic(err)
	}
}
