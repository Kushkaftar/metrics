package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
	"net/http"
)

func (h *Handler) getCounters(c *gin.Context) {

	d := models.Domain{
		ID:   1,
		Name: "domain1.com",
	}
	//h.lg.Info("1 - handler getCounters")
	counters, err := h.services.Counter.GetCounters(d)
	if err != nil {
		str := "something broken, could not get counters"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}
	c.JSON(http.StatusOK, counters)
}
