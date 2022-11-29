package promo

import (
	"go.uber.org/zap"
	"metrics/internal/models"
)

func (p *Promo) GetAllDomains() ([]models.Domain, error) {
	var domains []models.Domain

	folders, err := getFolder(p.path)
	if err != nil {
		p.lg.Error("Promo directory not found",
			zap.Error(err))
		return nil, err
	}

	for _, folder := range folders {
		domain := models.Domain{
			Name: folder}

		domains = append(domains, domain)
	}

	return domains, nil
}
