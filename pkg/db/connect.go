package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"metrics/pkg/config"
)

const (
	domainTable   = "domain"
	labelsTable   = "labels"
	countersTable = "counters"
)

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
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
