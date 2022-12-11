package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metrics/internal/models"
)

type LabelDB struct {
	db *sqlx.DB
	lg *zap.Logger
}

//todo refactor

func (db LabelDB) CreateLabel(label *models.Label) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (domain_id, metric_name, metric_id) values ($1, $2, $3) RETURNING id", labelsTable)

	row := db.db.QueryRow(query, label.DomainID, label.MetricName, label.MetricID)
	if err := row.Scan(&id); err != nil {
		db.lg.Error("CreateLabel",
			zap.Error(err))
		return err
	}

	label.ID = id

	return nil
}

//func (l LabelDB) GetLabelsInIdDomain(domain *models.Domain) ([]models.Label, error) {
//	var labels []models.Label
//
//	query := fmt.Sprintf("SELECT * FROM %s WHERE domain_id=$1", labelsTable)
//	if err := l.db.Select(&labels, query, domain.ID); err != nil {
//		l.lg.Error("GetAllLabels",
//			zap.Error(err))
//
//		return nil, err
//	}
//
//	return labels, nil
//}

func (db LabelDB) GetLabelInDomainID(domainID int) ([]models.Label, error) {
	var labels []models.Label

	query := fmt.Sprintf("SELECT * FROM %s WHERE domain_id=$1", labelsTable)
	if err := db.db.Select(&labels, query, domainID); err != nil {
		db.lg.Error("GetLabelInDomainID",

			zap.Int("domainID", domainID),
			zap.Error(err))
		return nil, err
	}

	return labels, nil
}

func (db LabelDB) GetLabelInID(id int) (*models.Label, error) {
	var label models.Label

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", labelsTable)
	if err := db.db.Get(&label, query, id); err != nil {
		db.lg.Error("GetLabelInID",

			zap.Int("id", id),
			zap.Error(err))
		return nil, err
	}

	return &label, nil
}

func (db LabelDB) GetLabelInName(label *models.Label) error {
	var labelDB models.Label

	query := fmt.Sprintf("SELECT * FROM %s WHERE metric_name=$1", labelsTable)
	if err := db.db.Get(&labelDB, query, label.MetricName); err != nil {
		db.lg.Error("GetLabelInName",

			zap.String("metric_name", label.MetricName),
			zap.Error(err))
		return err
	}

	label.ID = labelDB.ID
	label.DomainID = labelDB.DomainID
	label.MetricID = labelDB.MetricID

	return nil
}

func (db LabelDB) GetAllLabels() ([]models.Label, error) {
	var labels []models.Label

	query := fmt.Sprintf("SELECT * FROM %s", labelsTable)
	if err := db.db.Select(&labels, query); err != nil {
		db.lg.Error("GetAllLabels",
			zap.Error(err))

		return nil, err
	}

	return labels, nil
}

func newLabelDB(db *sqlx.DB, lg *zap.Logger) *LabelDB {
	return &LabelDB{
		db: db,
		lg: lg,
	}
}
