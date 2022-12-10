package db

import (
	"fmt"
	"metrics/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	domainTable   = "domain"
	labelsTable   = "labels"
	countersTable = "counters"
)

func NewPostgresDB(cfg models.DB) (*sqlx.DB, error) {
	dbConnect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", dbConnect)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
