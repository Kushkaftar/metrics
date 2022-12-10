package service

import (
	"metrics/internal/models"
	"metrics/pkg/db"
	"metrics/pkg/promo"

	"go.uber.org/zap"
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

func (srv *DomainSRV) GetAllDomains() ([]models.Domain, error) {
	// получем домены из репозитория
	promoDomains, err := srv.promo.GetAllDomains()
	if err != nil {
		return nil, err
	}

	// получаем домены из БД
	dbDomains, err := srv.db.GetAllDomains()
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
			if err := srv.db.CreateDomain(&domain); err != nil {
				// todo ???
				continue
			}
		}
	}

	// если есть домены, устанавливаем стутус "удален"
	if len(ignoreDB) != 0 {
		for _, domain := range ignoreDB {
			domain.Status = statusDelete
			if err := srv.db.UpdateStatus(&domain); err != nil {
				// todo ???
				continue
			}
		}
	}

	// повторно запрашиваем домены из БД
	dbDomains, err = srv.db.GetAllDomains()
	if err != nil {
		return nil, err
	}

	return dbDomains, nil
}

func (srv *DomainSRV) SetStatus(domain models.Domain) error {

	// устанавливаем статус
	if err := srv.db.UpdateStatus(&domain); err != nil {
		return err
	}

	return nil
}

func (srv *DomainSRV) Run() error {
	// получаем домены из БД
	dbDomains, err := srv.db.GetAllDomains()
	if err != nil {
		return err
	}

	// выбираем домены обхода
	for _, domain := range dbDomains {
		if domain.Status == statusWatch {
			// TODO: дописать логику
		}
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
