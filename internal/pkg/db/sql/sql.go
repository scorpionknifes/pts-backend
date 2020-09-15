package db

import (
	"fmt"
	"log"
	"os"

	"github.com/scorpionknifes/pts-backend/graph/model"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Db - gorm db
var Db *gorm.DB

// InitDB setup gorm
func InitDB() {
	var err error

	server := os.Getenv("SQL_SERVER")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")
	database := os.Getenv("SQL_DATABASE")

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	log.Println(dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migration to create tables for Story, User,  schema
	db.AutoMigrate(&model.Story{}, &model.Turn{}, &model.User{})
	Db = db
}

// Build connection string
