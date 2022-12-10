package service

import (
	"errors"
	"fmt"
	"metrics/internal/labelService"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"

	"go.uber.org/zap"
)

type labelSRV struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (srv *labelSRV) AddLabels(domain models.Domain) ([]models.Label, error) {
	var labels []models.Label

	// получаем домен из БД и берем его id
	domainDB, err := srv.db.Domain.GetDomain(domain.Name)
	if err != nil {
		return nil, err
	}

	// получаем метки домена
	labelsPromo, err := srv.promo.GetAllLabels(domain)
	if err != nil {
		return nil, err
	}

	// инитим пакет с логикой заведения/запроса метки
	ls := labelService.NewLabelService(srv.lg, srv.metrics, srv.db)

	// перебираем полученные метки из репозитория промо
	for _, labelPromo := range labelsPromo {
		labelPromo.DomainID = domainDB.ID

		// передаем метку для запроса/заведения в бд и метрике, если есть ошибки просто их логгируем
		check, err := ls.CheckLabel(&labelPromo)
		if err != nil {
			srv.lg.Error("error check label",
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
		lg:      lg,
		db:      db,
		promo:   promo,
		metrics: ym,
	}
}
