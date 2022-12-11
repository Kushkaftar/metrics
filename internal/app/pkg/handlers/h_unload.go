package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) unloadAllNewCounters(c *gin.Context) {

	paths, err := h.services.Counter.UnloadAllNewCounters()
	if err != nil {
		if err.Error() == "no new counters" {
			newErrorResponse(c, http.StatusBadRequest, "no new counters")
			return
		}
		str := "something broken"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	if len(paths) == 0 {
		str := "something broken, return null counters"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	for _, path := range paths {
		c.File(path)
	}
}

func (h *Handler) unload(c *gin.Context) {
	//var counters []models.Counter

	files, err := h.services.Counter.UnloadAllNewCounters()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	for _, file := range files {
		c.File(file)
	}
	return
}
