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

func (c CounterDB) CreateCounter(counter *models.Counter) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (metric_name, metric_id, label_id) values ($1, $2, $3) RETURNING id;", countersTable)

	row := c.db.QueryRow(query, counter.MetricName, counter.MetricID, counter.LabelID)
	if err := row.Scan(&id); err != nil {
		c.lg.Error("CreateCounter",
			zap.Error(err))
		return err
	}
	counter.ID = id
	return nil
}

func (c CounterDB) GetCounter(counter *models.Counter) error {

	query := fmt.Sprintf("SELECT * FROM %s WHERE metric_name=$1", countersTable)
	if err := c.db.Get(counter, query, counter.MetricName); err != nil {
		c.lg.Error("GetCounter",
			zap.Error(err))
		return err
	}

	return nil
}

func newCounterDB(db *sqlx.DB, lg *zap.Logger) *CounterDB {
	return &CounterDB{
		db: db,
		lg: lg,
	}
}
