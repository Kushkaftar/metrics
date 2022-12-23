package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metrics/internal/models"
)

type DomainDB struct {
	db *sqlx.DB
	lg *zap.Logger
}

// todo refactor

func (db DomainDB) CreateDomain(domain *models.Domain) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, status) values ($1, $2) RETURNING id", domainTable)

	row := db.db.QueryRow(query, domain.Name, domain.Status)
	if err := row.Scan(&id); err != nil {
		db.lg.Error("CreateDomain",
			zap.Error(err))
		return err
	}

	return nil
}

func (db DomainDB) GetDomain(domainName string) (*models.Domain, error) {
	var domain models.Domain

	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", domainTable)
	if err := db.db.Get(&domain, query, domainName); err != nil {
		db.lg.Error("GetDomain",
			zap.Error(err))
		return nil, err
	}

	return &domain, nil
}

func (db DomainDB) GetAllDomains() ([]models.Domain, error) {
	var domains []models.Domain

	query := fmt.Sprintf("SELECT * FROM %s", domainTable)
	if err := db.db.Select(&domains, query); err != nil {
		db.lg.Error("GetAllLabels",
			zap.Error(err))
		return nil, err
	}

	return domains, nil
}

func (db DomainDB) UpdateStatus(domain *models.Domain) error {

	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2;", domainTable)

	res, err := db.db.Exec(query, domain.Status, domain.ID)
	if err != nil {
		db.lg.Error("UpdateStatus db.Exec",
			zap.Error(err))
		return err
	}

	// todo del?
	count, err := res.RowsAffected()
	if err != nil {
		db.lg.Error("UpdateStatus RowsAffected",
			zap.Error(err))
		return err
	}

	if count != 1 {
		db.lg.Error("UpdateStatus  id out of range")
		return errors.New("id out of range")
	}

	return nil
}

func NewDomainDB(db *sqlx.DB, lg *zap.Logger) *DomainDB {
	return &DomainDB{
		db: db,
		lg: lg,
	}
}
