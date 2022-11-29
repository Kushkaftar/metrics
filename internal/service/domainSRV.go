package service

import (
	"go.uber.org/zap"
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/promo"
)

const (
	statusNew = iota
	statusWatch
	statusIgnore
	statusDelete
)

type DomainSRV struct {
	lg    *zap.Logger
	db    db.Domain
	promo *promo.Promo
}

func (d *DomainSRV) GetAllDomains() ([]models.Domain, error) {
	// получем домены из репозитория
	promoDomains, err := d.promo.GetAllDomains()
	if err != nil {
		return nil, err
	}

	// получаем домены из БД
	dbDomains, err := d.db.GetAllDomains()
	if err != nil {
		return nil, err
	}

	// сравниваем домены, возврящаем массивы для добавления в БД и установки флага "удален"
	addToDB, ignoreDB := compare(promoDomains, dbDomains)

	// если массивы пустые взвращаем домены из БД
	if len(addToDB) == 0 && len(ignoreDB) != 0 {
		return dbDomains, nil
	}

	//если есть домены, добавляем в БД
	if len(addToDB) != 0 {
		for _, domain := range addToDB {
			domain.Status = statusNew
			if err := d.db.CreateDomain(&domain); err != nil {
				// todo ???
				continue
			}
		}
	}

	// если есть домены, устанавливаем стутус "удален"
	if len(ignoreDB) != 0 {
		for _, domain := range ignoreDB {
			domain.Status = statusDelete
			if err := d.db.UpdateStatus(&domain); err != nil {
				// todo ???
				continue
			}
		}
	}

	// повторно запрашиваем домены из БД
	dbDomains, err = d.db.GetAllDomains()
	if err != nil {
		return nil, err
	}

	return dbDomains, nil
}

func (d *DomainSRV) SetStatus(domain models.Domain) error {

	// устанавливаем статус
	if err := d.db.UpdateStatus(&domain); err != nil {
		return err
	}

	return nil
}

func newDomainSRV(lg *zap.Logger, db db.Domain, promo *promo.Promo) *DomainSRV {
	return &DomainSRV{
		lg:    lg,
		db:    db,
		promo: promo,
	}
}
