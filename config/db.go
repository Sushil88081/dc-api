package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const DBURL = "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	} else {
		logrus.Info("Database connected  successfully")
	}
	return db
}
