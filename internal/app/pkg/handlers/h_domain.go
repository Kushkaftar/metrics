package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
	"net/http"
)

func (h *Handler) getAllDomains(c *gin.Context) {
	domains, err := h.services.Domain.GetAllDomains()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"domains": domains,
	})
}

func (h *Handler) setStatus(c *gin.Context) {
	var domain models.Domain

	if err := c.BindJSON(&domain); err != nil {
		str := "transmitted data is not valid"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	if !(domain.Status == 1 || domain.Status == 2) {
		str := "status is 1 or 2"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	if err := h.services.SetStatus(domain); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) change(c *gin.Context) {}

func (h *Handler) delete(c *gin.Context) {
	domain := models.NewDomain()

	if err := c.BindJSON(&domain); err != nil {
		str := "transmitted data is not valid"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	if domain.Status != 2 {
		str := "status not equal 2"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	if err := h.services.Domain.Delete(domain); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
