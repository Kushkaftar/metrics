package service

import (
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/app/service/counterService"
	"metrics/internal/app/service/labelService"
	"metrics/internal/models"
)

const (
	statusNew = iota
	statusWatch
	statusIgnore
	statusDelete
)

type DomainSRV struct {
	lg    *zap.Logger
	db    *db.DB
	promo *promo.Promo
	ym    *metrics.Metrics
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

// Run todo refactor
func (srv *DomainSRV) Run() error {

	// получаем домены из БД
	dbDomains, err := srv.db.GetAllDomains()
	if err != nil {
		return err
	}

	checkLabels := labelService.NewLabelService(srv.lg, srv.ym, srv.db, srv.promo)
	checkCounters := counterService.NewCounterService(srv.lg, srv.db, srv.promo, srv.ym)

	// выбираем домены обхода
	for _, domain := range dbDomains {
		if domain.Status == statusWatch {

			// получаем метки домена
			labels, err := checkLabels.CheckLabels(domain)
			if err != nil {
				// todo: add log error
				continue
			}

			//  проверяем что они есть
			if len(labels) != 0 {

				for _, label := range labels {

					counters, err := srv.promo.GetPromoUrls(&label)
					if err != nil {
						// todo: add log error
						continue
					}

					for _, counter := range counters {

						counter.LabelID = label.ID
						_, err := checkCounters.CheckCounter(&counter)
						if err != nil {
							// todo: add log error
							continue
						}

					}
				}
			}

		}
	}

	return nil
}

func newDomainSRV(lg *zap.Logger, db *db.DB, promo *promo.Promo, ym *metrics.Metrics) *DomainSRV {
	return &DomainSRV{
		lg:    lg,
		db:    db,
		promo: promo,
		ym:    ym,
	}
}
