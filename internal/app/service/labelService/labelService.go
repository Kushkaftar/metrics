package labelService

import (
	"go.uber.org/zap"
	"log"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/models"
)

type LabelService struct {
	lg    *zap.Logger
	ym    *metrics.Metrics
	db    *db.DB
	promo *promo.Promo
}

// todo: реализовать логгирование метки в случае ошибки

func (s *LabelService) CheckLabel(label *models.Label) (bool, error) {
	// проверяем есть ли метка в бд
	checkDB, err := s.checkLabelDB(label)
	if err != nil {
		return false, err
	}
	// если метка есть выходим
	if checkDB {
		return true, nil
	}

	// проверяем заведена ли метка в метрике
	checkMetric, err := s.checkLabelMetric(label)
	if err != nil {
		return false, err
	}

	// если заведена, добаляем в базу
	if checkMetric {
		// если все ок выходим
		if _, err = s.createLabelDB(label); err != nil {
			return false, err
		}

		return true, nil
	}

	// todo del log
	log.Printf("labelPromo - %+v", label)
	// создаем метку в метрике
	if _, err = s.createLabelMetric(label); err != nil {
		return false, err
	}

	// добавляем метку в базу
	if _, err = s.createLabelDB(label); err != nil {
		return false, err
	}

	return true, nil

}

func (s *LabelService) checkLabelDB(label *models.Label) (bool, error) {

	if err := s.db.GetLabelInName(label); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (s *LabelService) createLabelDB(label *models.Label) (bool, error) {
	if err := s.db.CreateLabel(label); err != nil {
		return false, err
	}

	return true, nil
}

func (s *LabelService) checkLabelMetric(label *models.Label) (bool, error) {
	labelsMetric, err := s.ym.GetAllLabels()
	if err != nil {
		return false, err
	}

	for _, l := range labelsMetric {
		if l.MetricName == label.MetricName {
			label.MetricID = l.MetricID

			return true, nil
		}
	}

	return false, nil
}

func (s *LabelService) createLabelMetric(label *models.Label) (bool, error) {
	if err := s.ym.CreateLabel(label); err != nil {
		return false, err
	}

	return true, nil
}

func NewLabelService(lg *zap.Logger, ym *metrics.Metrics, db *db.DB, promo *promo.Promo) *LabelService {
	return &LabelService{
		lg:    lg,
		ym:    ym,
		db:    db,
		promo: promo,
	}
}
