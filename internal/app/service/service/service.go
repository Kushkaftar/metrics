package service

import (
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/models"
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

type Download interface {
	GetLabelList(date string) ([]models.Label, error)
	DownloadFile(id int) (string, error)
}

type Service struct {
	Domain
	Label
	Counter
	Download
}

func NewService(lg *zap.Logger, ym *metrics.Metrics, db *db.DB, promo *promo.Promo) *Service {
	return &Service{
		Domain:   newDomainSRV(lg, db, promo, ym),
		Label:    newLabelSRV(lg, db, promo, ym),
		Counter:  newCounterSRV(lg, db, promo, ym),
		Download: newDownloadSRV(lg, db, promo, ym),
	}
}
