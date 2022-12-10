package service

import (
	"errors"
	"go.uber.org/zap"
	"metrics/internal/counterService"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"
	"metrics/pkg/utils/unload"
)

type CounterSRV struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (srv *CounterSRV) UnloadAllNewCounters() ([]string, error) {
	var paths []string

	// запрашиваем метки созданных счетчиков
	labelsNewCounter, err := srv.db.Counter.GetLabelInNewCounters()
	if err != nil {
		return nil, err
	}

	// если новых счетчиков нет выходим
	if len(labelsNewCounter) == 0 {
		return nil, errors.New("no new counters")
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
	srv.lg.Info("2 - SERVICE GetCounters, crete slice domainCounters")
	check := counterService.NewCounterService(srv.lg, srv.db, srv.promo, srv.metrics)
	srv.lg.Info("3 - SERVICE GetCounters, init check counter")
	// получаем метки домена
	labels, err := srv.db.Label.GetLabelInDomainID(domain.ID)
	srv.lg.Info("4 - SERVICE GetCounters, get labels")

	if err != nil {
		return nil, err
	}

	for _, label := range labels {
		srv.lg.Info("5 - SERVICE GetCounters, range labels")

		counters, err := srv.promo.GetPromoUrls(&label)
		srv.lg.Info("6 - SERVICE GetCounters, gel all lands in promo")
		if err != nil {
			return nil, err
		}

		for _, counter := range counters {
			counter.LabelID = label.ID

			if _, err := check.CheckCounter(&counter); err != nil {
				srv.lg.Error("error check counter",
					zap.Error(err))
			}
			srv.lg.Info("7 - SERVICE GetCounters, check counter")
			//log.Printf("counter - %+v", counter)
			domainCounters = append(domainCounters, counter)
		}

	}
	srv.lg.Info("8 - SERVICE GetCounters, return all counters")
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
