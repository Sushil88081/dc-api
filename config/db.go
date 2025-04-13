package config

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DBURL = "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := DBURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	logrus.Info("✅ Connected to PostgreSQL database successfully.")
	DB = db
	return db
}
