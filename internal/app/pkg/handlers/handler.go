package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metrics/internal/app/service/service"
)

type Handler struct {
	services *service.Service
	lg       *zap.Logger
}

func NewHandler(s *service.Service, lg *zap.Logger) *Handler {
	return &Handler{
		services: s,
		lg:       lg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/unload", h.unload)
		api.GET("/run", h.run)

		domain := api.Group("domain")
		{
			domain.GET("", h.getAllDomains)
			domain.POST("/set_status", h.setStatus)
			domain.POST("/check_labels", h.addLabels)

			labels := domain.Group("/:id/labels")
			{
				labels.GET("", h.domainLabels)
			}

		}
		counters := api.Group("counters")
		{
			counters.POST("/check", h.checkCounters)
			counters.GET("", h.getCounters)
		}
	}

	return router
}
