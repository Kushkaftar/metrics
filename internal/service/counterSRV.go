package service

import (
	"go.uber.org/zap"
	"metrics/internal/counterService"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"
)

type CounterSRV struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (c *CounterSRV) GetCounters(domain models.Domain) ([]models.Counter, error) {
	var domainCounters []models.Counter
	c.lg.Info("2 - SERVICE GetCounters, crete slice domainCounters")
	check := counterService.NewCounterService(c.lg, c.db, c.promo, c.metrics)
	c.lg.Info("3 - SERVICE GetCounters, init check counter")
	// получаем метки домена
	labels, err := c.db.Label.GetLabelInDomainID(domain.ID)
	c.lg.Info("4 - SERVICE GetCounters, get labels")

	if err != nil {
		return nil, err
	}

	for _, label := range labels {
		c.lg.Info("5 - SERVICE GetCounters, range labels")

		counters, err := c.promo.GetPromoUrls(&label)
		c.lg.Info("6 - SERVICE GetCounters, gel all lands in promo")
		if err != nil {
			return nil, err
		}

		for _, counter := range counters {
			counter.LabelID = label.ID

			if _, err := check.CheckCounter(&counter); err != nil {
				c.lg.Error("error check counter",
					zap.Error(err))
			}
			c.lg.Info("7 - SERVICE GetCounters, check counter")
			//log.Printf("counter - %+v", counter)
			domainCounters = append(domainCounters, counter)
		}

	}
	c.lg.Info("8 - SERVICE GetCounters, return all counters")
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
