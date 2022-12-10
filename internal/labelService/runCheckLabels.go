package labelService

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"metrics/internal/models"
)

func (s *LabelService) CheckLabels(domain models.Domain) ([]models.Label, error) {
	var labels []models.Label

	// получаем домен из БД и берем его id
	domainDB, err := s.db.Domain.GetDomain(domain.Name)
	if err != nil {
		return nil, err
	}

	// получаем метки домена
	labelsPromo, err := s.promo.GetAllLabels(domain)
	if err != nil {
		return nil, err
	}

	// перебираем полученные метки из репозитория промо
	for _, labelPromo := range labelsPromo {
		labelPromo.DomainID = domainDB.ID

		// передаем метку для запроса/заведения в бд и метрике, если есть ошибки просто их логгируем
		check, err := s.CheckLabel(&labelPromo)
		if err != nil {
			s.lg.Error("error check label",
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
		return nil, errors.New(strErr)
	}

	return labels, nil
}
