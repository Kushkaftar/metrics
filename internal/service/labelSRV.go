package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"metrics/internal/labelService"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"
)

type labelSRV struct {
	lg      *zap.Logger
	db      db.Label
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (l *labelSRV) GetLabels(domain models.Domain) ([]models.Label, error) {
	var labels []models.Label

	// получаем метки домена
	labelsPromo, err := l.promo.GetAllLabels(domain)
	if err != nil {
		return nil, err
	}

	// инитим пакет с логикой заведения/запроса метки
	ls := labelService.NewLabelService(l.lg, l.metrics, l.db)

	// перебираем полученные метки из репозитория промо
	for _, labelPromo := range labelsPromo {

		// передаем метку для запроса/заведения в бд и метрике, если есть ошибки просто их логгируем
		check, err := ls.CheckLabel(&labelPromo)
		if err != nil {
			l.lg.Error("error check label",
				zap.Error(err))
		}

		// если все ок, добавляем метку в массив
		if check {
			labels = append(labels, labelPromo)
		}
	}

	// проверяем срез на длинну, если она равна нулю, считаем что метки не удалось добавить
	if len(labels) == 0 {
		strErr := fmt.Sprintf("Failed to add labels: %+v", labelsPromo)
		errStr := errors.New(strErr)
		return nil, errStr
	}

	return labels, nil
}

func newLabelSRV(lg *zap.Logger, db db.Label, promo *promo.Promo, ym *metrics.Metrics) *labelSRV {
	return &labelSRV{
		lg:      lg,
		db:      db,
		promo:   promo,
		metrics: ym,
	}
}
