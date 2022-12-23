package service

import (
	"errors"
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/models"
	"metrics/pkg/unload"
)

type DownloadSRV struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (srv *DownloadSRV) GetLabelList(date string) ([]models.Label, error) {
	// запрашиваем счетчики созданные в переданную дату, получаем id меток
	labelsID, err := srv.db.Counter.GetLabelInNewCounters(date)
	if err != nil {
		srv.lg.Error("Service, GetLabelList",
			zap.Error(err))
		return nil, errors.New("not found")
	}

	// запрашиваем метки по переданным id
	var labels []models.Label

	for _, labelID := range labelsID {
		label, err := srv.db.Label.GetLabelInID(labelID)
		if err != nil {
			srv.lg.Error("Service, GetLabelList",
				zap.Error(err))
			continue
		}
		if label != nil {
			labels = append(labels, *label)
		}
	}

	// возвращаем массив меток
	if len(labels) == 0 {
		return nil, errors.New("not found")
	}

	return labels, nil
}

func (srv *DownloadSRV) DownloadFile(id int) (string, error) {

	// запрашиваем метку id
	label, err := srv.db.Label.GetLabelInID(id)
	if err != nil {
		srv.lg.Error("Service, DownloadFile",
			zap.Error(err))
		return "", err
	}

	// получаем счетчики метки
	counters, err := srv.db.Counter.GetCountersLabel(label.ID)
	if err != nil {
		srv.lg.Error("Service, DownloadFile",
			zap.Error(err))
		return "", err
	}

	filePath, err := unload.Unload(label.MetricName, counters)
	if err != nil {
		srv.lg.Error("Service, DownloadFile",
			zap.Error(err))
		return "", err
	}

	return filePath, nil
}

func newDownloadSRV(lg *zap.Logger, db *db.DB, promo *promo.Promo, ym *metrics.Metrics) *DownloadSRV {
	return &DownloadSRV{
		lg:      lg,
		db:      db,
		promo:   promo,
		metrics: ym,
	}
}
