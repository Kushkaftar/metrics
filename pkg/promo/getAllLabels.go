package promo

import (
	"fmt"
	"go.uber.org/zap"
	"metrics/internal/models"
)

func (p *Promo) GetAllLabels(domain models.Domain) ([]models.Label, error) {
	var labels []models.Label

	path := fmt.Sprintf("%s%s", p.path, domain.Name)

	folders, err := getFolder(path)
	if err != nil {
		p.lg.Error("Promo directory not found",
			zap.Error(err))
		return nil, err
	}

	for _, folder := range folders {
		labelName := fmt.Sprintf("%s_%s", domain.Name, folder)
		label := models.Label{
			DomainID:   domain.ID,
			MetricName: labelName,
		}

		labels = append(labels, label)
	}

	return labels, nil

}
