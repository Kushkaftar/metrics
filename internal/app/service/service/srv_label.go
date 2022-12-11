package service

import (
	"errors"
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/app/service/labelService"
	"metrics/internal/models"
)

type labelSRV struct {
	lg    *zap.Logger
	db    *db.DB
	promo *promo.Promo
	ym    *metrics.Metrics
}

func (srv *labelSRV) AddLabels(domain models.Domain) ([]models.Label, error) {

	ls := labelService.NewLabelService(srv.lg, srv.ym, srv.db, srv.promo)

	labels, err := ls.CheckLabels(domain)
	if err != nil {
		return nil, err
	}

	return labels, nil
}

func (srv *labelSRV) GetLabels(domain models.Domain) ([]models.Label, error) {
	labels, err := srv.db.GetLabelInDomainID(domain.ID)
	if err != nil {
		return nil, err
	}

	if len(labels) == 0 {
		return nil, errors.New("labels not found")
	}

	return labels, nil
}

func newLabelSRV(lg *zap.Logger, db *db.DB, promo *promo.Promo, ym *metrics.Metrics) *labelSRV {
	return &labelSRV{
		lg:    lg,
		db:    db,
		promo: promo,
		ym:    ym,
	}
}
