package service

import (
	"metrics/internal/models"
)

// todo refactor
func compare(promoDomains []models.Domain, dbDomains []models.Domain) (db []models.Domain, del []models.Domain) {
	arraySeparation := make(map[string]string)

	for _, item := range promoDomains {
		arraySeparation[item.Name] = "add"
	}
	for _, item := range dbDomains {
		if arraySeparation[item.Name] == "add" {
			arraySeparation[item.Name] = "ignore"
		} else {
			arraySeparation[item.Name] = "del"
		}
	}

	var (
		toDel []models.Domain
		toDB  []models.Domain
	)
	for key, value := range arraySeparation {
		switch value {
		case "add":
			for _, domain := range promoDomains {
				if domain.Name == key {
					toDB = append(toDB, domain)
				}
			}

		case "del":
			for _, domain := range dbDomains {
				if domain.Name == key {
					toDel = append(toDel, domain)
				}
			}
		default:

		}
	}
	return toDB, toDel
}
