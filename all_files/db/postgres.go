package db

import (
	"log"
	"time"

	"taskmanage/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func ConnectPostgres(dsn string) *gorm.DB {
	logger.Ok, _ = logger.GetLogger("database")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logger.Ok.Println("Could not gyeee registration: ", err.Error())
		log.Fatal("Error connecting to database")

	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to configure SQL DB: ", err)
	}

	// Set connection pool configurations
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}
