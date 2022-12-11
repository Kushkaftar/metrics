package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) run(c *gin.Context) {
	if err := h.services.Domain.Run(); err != nil {
		str := "some err"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
