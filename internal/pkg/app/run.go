package app

import (
	"go.uber.org/zap"
	"metrics/internal/app/pkg/db"
	"metrics/internal/app/pkg/handlers"
	"metrics/internal/app/pkg/metrics"
	"metrics/internal/app/pkg/promo"
	"metrics/internal/app/service/service"
	"metrics/internal/models"
	"metrics/pkg/logging"

	_ "github.com/lib/pq"
)

func Run(c *models.Config) error {
	// LOGGER
	lg := logging.NewLogger(c)

	// DB POSTGRES
	connect, err := db.NewPostgresDB(c.DB)
	if err != nil {
		lg.Error("postgres fail starting",
			zap.Error(err))
		return err
	}
	dbp := db.NewDB(connect, lg)

	// PROMO
	p := promo.NewPromo(&c.Promo, lg)

	// Y METRIC
	ym := metrics.NewMetrics(c, lg)

	// SERVICE
	src := service.NewService(lg, ym, dbp, p)

	// HANDLERS
	handler := handlers.NewHandler(src, lg)

	// SERVER
	srv := new(server)
	if err = srv.serverRun(c.SRV.Port, handler.InitRoutes()); err != nil {
		lg.Error("server fail starting",
			zap.Error(err))
		return err
	}

	return nil
}
