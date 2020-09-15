package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var server = os.Getenv("SQL_SERVER")
var port = os.Getenv("SQL_PORT")
var user = os.Getenv("SQL_USER")
var password = os.Getenv("SQL_PASSWORD")
var database = os.Getenv("SQL_DATABASE")

// InitDB setup gorm
func InitDB() {
	var err error
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migration to create tables for Order and Item schema
	db.AutoMigrate()
}

// Build connection string
