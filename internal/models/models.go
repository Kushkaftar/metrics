package models

type Domain struct {
	ID     int    `json:"id"`
	Name   string `db:"name" json:"name"`
	Status int    `db:"status" json:"status"`
}

type Label struct {
	ID         int    `db:"id" json:"id"`
	DomainID   int    `db:"domain_id" json:"domain_id"`
	MetricName string `db:"metric_name"`
	MetricID   int    `db:"metric_id"`
}

type Counter struct {
	ID         int    `json:"id"`
	MetricName string `db:"metric_name"`
	MetricID   int    `db:"metric_id"`
	LabelID    int    `db:"label_id"`
}
