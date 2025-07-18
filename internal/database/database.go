package database

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/gotus/internal/service/log"
	"os"
	"time"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	))

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetConnMaxIdleTime(time.Second * 5)

	return db
}
