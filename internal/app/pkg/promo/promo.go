package promo

import (
	"metrics/internal/models"

	"go.uber.org/zap"
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

func NewPromo(c *models.Promo, lg *zap.Logger) *Promo {
	return &Promo{
		lg:   lg,
		path: c.Path,
	}
}
