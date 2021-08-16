package database

import (
	"clean-architecture/infrastructure/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	*gorm.DB
}

var singleton = newGorm()

func newGorm() *Gorm {
	DB_HOST := "localhost"
	DB_PORT := "5432"
	DB_USER := "clean"
	DB_NAME := "clean"
	DB_PASS := "clean"

	// https://stackoverflow.com/questions/21959148/ssl-is-not-enabled-on-the-server
	CONNEXT := "host=" + DB_HOST + " port=" + DB_PORT + " user=" + DB_USER + " dbname=" + DB_NAME + " password=" + DB_PASS + " sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(CONNEXT), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}

	if !db.Migrator().HasTable("users") {
		panic("not exists users table")
	} else {
		logger.Debug("exists users table")
	}

	if !db.Migrator().HasTable("tasks") {
		panic("not exists tasks table")
	} else {
		logger.Debug("exists tasks table")
	}

	return &Gorm{db}
}

func GormConnect() *Gorm {
	return singleton
}
