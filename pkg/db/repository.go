package db

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metrics/internal/models"
)

type Domain interface {
	CreateDomain(domain *models.Domain) error
	GetDomain(domainName string) (*models.Domain, error)
	GetAllDomains() ([]models.Domain, error)
	UpdateStatus(domain *models.Domain) error
}

type Label interface {
	CreateLabel(label *models.Label) error
	GetAllLabels() ([]models.Label, error)
	GetLabelInID(id int) (*models.Label, error)
	GetLabelInName(label *models.Label) error
	GetLabelInDomainID(domainID int) ([]models.Label, error)
}

type Counter interface {
	CreateCounter(counter *models.Counter) error
	GetCounter(counter *models.Counter) error
	GetLabelInNewCounters() ([]int, error)
	GetCountersLabel(LabelID int) ([]models.Counter, error)
}

type DB struct {
	Domain
	Label
	Counter
}

func NewDB(db *sqlx.DB, lg *zap.Logger) *DB {
	return &DB{
		Domain:  NewDomainDB(db, lg),
		Label:   newLabelDB(db, lg),
		Counter: newCounterDB(db, lg),
	}
}
