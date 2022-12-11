package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
	"net/http"
)

func (h *Handler) checkCounters(c *gin.Context) {
	var domain models.Domain

	if err := c.BindJSON(&domain); err != nil {
		str := "transmitted json is not valid"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	h.lg.Info("1 - handler getCounters")
	counters, err := h.services.Counter.GetCounters(domain)
	if err != nil {
		str := "something broken, could not get counters"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	// todo add in service check label
	if len(counters) == 0 {
		str := "counters not found"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	c.JSON(http.StatusOK, counters)
}

func (h *Handler) getCounters(c *gin.Context) {

}
