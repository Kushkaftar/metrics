package handlers

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/models"
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
	var counters []models.Counter
	c1 := models.Counter{
		MetricName: "test1",
		MetricID:   12,
	}
	c2 := models.Counter{
		MetricName: "test2",
		MetricID:   1223465,
	}
	counters = append(counters, c1, c2)

	//if err := unload.Unload("test_counter", counters); err != nil {
	//	log.Printf("err - %s", err)
	//	return
	//}
	c.File("./tmp/counter.js")
}
