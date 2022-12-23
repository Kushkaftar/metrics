package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metrics/internal/models"
)

type CounterDB struct {
	db *sqlx.DB
	lg *zap.Logger
}

func (db CounterDB) CreateCounter(counter *models.Counter) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (metric_name, metric_id, label_id) values ($1, $2, $3) RETURNING id;", countersTable)

	row := db.db.QueryRow(query, counter.MetricName, counter.MetricID, counter.LabelID)
	if err := row.Scan(&id); err != nil {
		db.lg.Error("CreateCounter",
			zap.Error(err))
		return err
	}
	counter.ID = id
	return nil
}

func (db CounterDB) GetCounter(counter *models.Counter) error {

	query := fmt.Sprintf("SELECT * FROM %s WHERE metric_name=$1", countersTable)
	if err := db.db.Get(counter, query, counter.MetricName); err != nil {
		db.lg.Error("GetCounter",
			zap.Error(err))
		return err
	}

	return nil
}

func (db CounterDB) GetLabelInNewCounters(date string) ([]int, error) {
	var labels []int

	query := fmt.Sprintf("SELECT DISTINCT label_id FROM %s WHERE created_at=$1", countersTable)

	//now := time.Now()
	//date := now.Format("2006-01-02")

	if err := db.db.Select(&labels, query, date); err != nil {
		db.lg.Error("GetLabelInNewCounters",
			zap.Error(err))
		return nil, err
	}

	return labels, nil
}

func (db CounterDB) GetCountersLabel(LabelID int) ([]models.Counter, error) {
	var counters []models.Counter

	query := fmt.Sprintf("SELECT * FROM %s WHERE label_id=$1", countersTable)
	if err := db.db.Select(&counters, query, LabelID); err != nil {
		db.lg.Error("GetCountersLabel",
			zap.Error(err))
		return nil, err
	}

	return counters, nil
}

func newCounterDB(db *sqlx.DB, lg *zap.Logger) *CounterDB {
	return &CounterDB{
		db: db,
		lg: lg,
	}
}
