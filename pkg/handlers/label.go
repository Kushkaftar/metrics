package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
	"net/http"
)

func (h *Handler) getLabels(c *gin.Context) {

	d := models.Domain{
		ID:   1,
		Name: "domain1.com",
	}
	labels, err := h.services.GetLabels(d)
	if err != nil {
		str := "something broken, could not get labels"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	c.JSON(http.StatusOK, labels)
}
