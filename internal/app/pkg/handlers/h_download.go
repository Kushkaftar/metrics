package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getListFileToDownload(c *gin.Context) {
	dateString := c.Query("date")

	_, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		h.lg.Error("err parse date",
			zap.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	labels, err := h.services.Download.GetLabelList(dateString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"labels": labels,
	})
}

func (h *Handler) downloadFile(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		str := "id is not a number"
		newErrorResponse(c, http.StatusBadRequest, str)
		return
	}

	downloadFile, err := h.services.Download.DownloadFile(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.File(downloadFile)
	// удаляем все файлы из папки для хранения файлов с кодом счетчиков
	// todo: add delete old files
}
