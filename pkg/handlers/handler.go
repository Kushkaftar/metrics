package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metrics/internal/service"
)

type Handler struct {
	services *service.Service
	lg       *zap.Logger
}

func NewHandler(s *service.Service, lg *zap.Logger) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		domain := api.Group("domain")
		{
			domain.GET("/", h.getAllDomains)
			domain.POST("/", h.setStatus)
		}
		label := api.Group("label")
		{
			label.POST("/", h.getLabels)
		}
		counters := api.Group("counters")
		{
			counters.POST("/", h.getCounters)
		}
	}

	router.Use(cors.Default())
	return router
}
