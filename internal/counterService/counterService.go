package counterService

import (
	"go.uber.org/zap"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/metrics"
	"metrics/pkg/promo"
)

type CounterService struct {
	lg      *zap.Logger
	db      *db.DB
	promo   *promo.Promo
	metrics *metrics.Metrics
}

func (s *CounterService) CheckCounter(counter *models.Counter) (bool, error) {
	checkDB, err := s.checkCounterDB(counter)
	if err != nil {
		return false, err
	}

	if checkDB {
		return true, nil
	}

	checkMetric, err := s.checkCounterMetric(counter)
	if err != nil {
		return false, err
	}

	if checkMetric {
		if _, err := s.createCounterDB(counter); err != nil {
			return false, err
		}

		return true, nil
	}

	if _, err := s.createCounterMetric(counter); err != nil {
		return false, err
	}

	if _, err := s.createCounterDB(counter); err != nil {
		return false, err
	}

	return true, nil
}

func (s *CounterService) checkCounterDB(counter *models.Counter) (bool, error) {
	if err := s.db.Counter.GetCounter(counter); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (s *CounterService) createCounterDB(counter *models.Counter) (bool, error) {
	if err := s.db.Counter.CreateCounter(counter); err != nil {
		return false, err
	}
	return true, nil
}

func (s *CounterService) checkCounterMetric(counter *models.Counter) (bool, error) {
	if err := s.metrics.GetCounter(counter); err != nil {
		if err.Error() == "counter not found" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *CounterService) createCounterMetric(counter *models.Counter) (bool, error) {

	if err := s.metrics.CreateCounter(counter); err != nil {
		return false, err
	}

	label, err := s.db.Label.GetLabelInID(counter.LabelID)
	if err != nil {
		return false, err
	}

	if err = s.metrics.SetLabelInCounter(counter, label); err != nil {
		return false, err
	}
	return true, nil
}

func NewCounterService(lg *zap.Logger, db *db.DB, promo *promo.Promo, ym *metrics.Metrics) *CounterService {
	return &CounterService{
		lg:      lg,
		db:      db,
		promo:   promo,
		metrics: ym,
	}
}
