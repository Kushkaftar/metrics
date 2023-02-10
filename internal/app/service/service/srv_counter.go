package service

import (
	"errors"
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/app/service/counterService"
	"metrics/internal/models"
	"metrics/pkg/errWrap"
	"metrics/pkg/unload"
	"time"
)

type CounterSRV struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

var (
	ErrorsNoNewCounter = errors.New("no new counters")
)

func (srv *CounterSRV) UnloadAllNewCounters() ([]string, error) {
	var paths []string

	now := time.Now()
	date := now.Format("2006-01-02")

	// запрашиваем метки созданных счетчиков
	labelsNewCounter, err := srv.db.Counter.GetLabelInNewCounters(date)
	if err != nil {
		return nil, err
	}

	// если новых счетчиков нет выходим
	if len(labelsNewCounter) == 0 {
		return nil, errWrap.Wrap("unload new counters", ErrorsNoNewCounter)
	}

	// удаляем все файлы из папки для хранения файлов с кодом счетчиков
	// todo: add delete old files

	// перебираем срез с метками
	for _, labelID := range labelsNewCounter {

		// запрашиваем в БД все счетчики метки
		counters, err := srv.db.Counter.GetCountersLabel(labelID)
		if err != nil {
			srv.lg.Error("error in GetCountersLabel",
				zap.Int("labelID", labelID),
				zap.Error(err))
		}

		// получаем метку из бд
		label, err := srv.db.Label.GetLabelInID(labelID)

		// !!! todo refactor
		if err == nil {

			// передаем счетчики и имя метки для записи файла
			fileName, err := unload.Unload(label.MetricName, counters)
			if err != nil {
				srv.lg.Error("error in write file",
					zap.Error(err))
			}

			// добавляем название файла в срез
			paths = append(paths, fileName)

		} else {
			srv.lg.Error("error in GetLabelInID",
				zap.Int("labelID", labelID),
				zap.Error(err))
		}

	}

	// возврящаем срез имен файлов
	return paths, nil
}

func (srv *CounterSRV) GetCounters(domain models.Domain) ([]models.Counter, error) {
	var domainCounters []models.Counter

	check := counterService.NewCounterService(srv.lg, srv.db, srv.promo, srv.metrics)

	// получаем метки домена
	labels, err := srv.db.Label.GetLabelInDomainID(domain.ID)

	if err != nil {
		return nil, err
	}

	for _, label := range labels {

		counters, err := srv.promo.GetPromoUrls(&label)
		if err != nil {
			return nil, err
		}

		for _, counter := range counters {
			counter.LabelID = label.ID

			if _, err := check.CheckCounter(&counter); err != nil {
				srv.lg.Error("error check counter",
					zap.Error(err))
			}

			//log.Printf("counter - %+v", counter)
			domainCounters = append(domainCounters, counter)
		}

	}

	// todo add GetCounters
	return domainCounters, nil
}

func newCounterSRV(lg *zap.Logger, db *db.DB, promo *promo.Promo, ym *metrics.Metrics) *CounterSRV {
	return &CounterSRV{
		lg:      lg,
		db:      db,
		promo:   promo,
		metrics: ym,
	}
}
