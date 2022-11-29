package metrics

import (
	"go.uber.org/zap"
	"metrics/internal/models"
	"metrics/pkg/client"
	"metrics/pkg/config"
	"metrics/pkg/utils/urlKit"
)

const (
	metricScheme = "https"
	metricHost   = "api-metrika.yandex.net"
)

type Metrics struct {
	lg     *zap.Logger
	url    *urlKit.Scheme
	client *client.Client
}

type LabelMethods interface {
	CreateLabel(label *models.Label) error
	GetAllLabels() ([]models.Label, error)
	SetLabelInCounter(counter *models.Counter, label *models.Label) error
}

type CounterMethods interface {
	GetCounter(counter *models.Counter) error
	CreateCounter(counter *models.Counter) error
	GetCounters(counterName string) ([]models.Counter, error)
	DelCounter(counter *models.Counter) error
}

func NewMetrics(c *config.Config, lg *zap.Logger) *Metrics {

	return &Metrics{
		lg:     lg,
		url:    urlKit.NewScheme(metricScheme, metricHost),
		client: client.NewClient(c.Metrics.Token),
	}
}
