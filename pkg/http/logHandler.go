package http

import (
	"logSaver/pkg/model"
	"logSaver/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	DB *store.DB
}

func (lh *LogHandler) CreateLog(context *gin.Context) {
	logData := model.Log{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	logData.Cost = 0.0
	err := lh.DB.LogRepository.Create(logData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}

func (lh *LogHandler) UpdateLog(context *gin.Context) {
	newLogData := model.Log{}
	if err := context.BindJSON(&newLogData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	existLogData, err := lh.DB.LogRepository.Read(&newLogData)
	if err != nil {
		return
	}

	updateExistLog(&existLogData, &newLogData)

	err = lh.DB.LogRepository.Update(existLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log updated"})
}

func updateExistLog(existLogData *model.Log, newLogData *model.Log) {
	existLogData.Status = newLogData.Status
	existLogData.StatusDelive = newLogData.StatusDelive
	existLogData.Cost = newLogData.Cost
	existLogData.Updated = newLogData.Updated
}
