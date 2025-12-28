package postgresql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBParams struct {
	Host                string
	Port                string
	Username            string
	Password            string
	Name                string
	SSLMode             string
	TimeZome            string
	MaxIdleConn         int
	MaxOpenConn         int
	MaxIdleInMinute     int
	MaxLifetimeInMinute int
}

func InitDbConnection(dbParams DBParams) *gorm.DB {
	connectionString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dsn := fmt.Sprintf(
		connectionString,
		dbParams.Host,
		dbParams.Username,
		dbParams.Password,
		dbParams.Name,
		dbParams.Port,
		dbParams.SSLMode,
		dbParams.TimeZome,
	)

	db, errors := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errors != nil {
		panic(errors.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get raw DB from GORM", err)
		panic(err.Error())
	}

	// Set connection pool configuration
	maxIdleConn := dbParams.MaxIdleConn
	maxOpenConn := dbParams.MaxOpenConn
	connMaxIdleTimeInMinute := dbParams.MaxIdleInMinute
	connMaxLifetimeInMinute := dbParams.MaxLifetimeInMinute

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(connMaxIdleTimeInMinute) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetimeInMinute) * time.Minute)

	return db
}
