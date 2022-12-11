package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) addLabels(c *gin.Context) {
	var domain models.Domain
	if err := c.BindJSON(&domain); err != nil {
		str := "transmitted json is not valid"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	labels, err := h.services.AddLabels(domain)
	if err != nil {
		str := "something broken, could not get labels"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	c.JSON(http.StatusOK, labels)
}

func (h *Handler) domainLabels(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		str := "id is not a number"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	domain := models.Domain{
		ID: id,
	}

	domains, err := h.services.Label.GetLabels(domain)
	if err != nil {
		str := "labels not found"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	c.JSON(http.StatusOK, domains)
}
