package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/luthfiybk/go-with-gin/src/helper"
)

const (
	host	 = "localhost"
	port	 = 5432
	user	 = "postgres"
	password = "root"
	dbName	 = "go_with_gin"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{}, &gorm.Config{})

	helper.ErrorPanic(err)

	return db
}
