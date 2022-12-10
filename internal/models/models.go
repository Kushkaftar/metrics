package models

import "time"

type Domain struct {
	ID     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Status int    `db:"status" json:"status"`
}

type Label struct {
	ID         int    `db:"id" json:"id"`
	DomainID   int    `db:"domain_id" json:"domain_id"`
	MetricName string `db:"metric_name" json:"metric_name"`
	MetricID   int    `db:"metric_id" json:"metricID"`
}

type Counter struct {
	ID         int       `db:"id" json:"id"`
	MetricName string    `db:"metric_name" json:"metric_name"`
	MetricID   int       `db:"metric_id" json:"metric_id"`
	LabelID    int       `db:"label_id" json:"label_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
