package service

import (
	"go.uber.org/zap"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"
)

// todo: refactor

type Domain interface {
	GetAllDomains() ([]models.Domain, error)
	SetStatus(models.Domain) error
	Run() error
}

type Label interface {
	AddLabels(models.Domain) ([]models.Label, error)
	GetLabels(domain models.Domain) ([]models.Label, error)
}

type Counter interface {
	GetCounters(domain models.Domain) ([]models.Counter, error)
	UnloadAllNewCounters() ([]string, error)
}

type Service struct {
	Domain
	Label
	Counter
}

func NewService(lg *zap.Logger, ym *metrics.Metrics, db *db.DB, promo *promo.Promo) *Service {
	return &Service{
		Domain:  newDomainSRV(lg, db.Domain, promo),
		Label:   newLabelSRV(lg, db, promo, ym),
		Counter: newCounterSRV(lg, db, promo, ym),
	}
}
