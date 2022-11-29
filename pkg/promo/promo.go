package promo

import (
	"go.uber.org/zap"
	"metrics/internal/models"
	"metrics/pkg/config"
)

type Promo struct {
	lg   *zap.Logger
	path string
}

type Methods interface {
	GetAllDomains() ([]models.Domain, error)
	GetAllLabels(models.Domain) ([]models.Label, error)
	GetPromoUrls(domain *models.Domain) ([]models.Counter, error)
}

func NewPromo(c *config.Promo, lg *zap.Logger) *Promo {
	return &Promo{
		lg:   lg,
		path: c.Path,
	}
}
