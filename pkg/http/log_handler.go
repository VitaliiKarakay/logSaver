package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"logSaver/pkg/model"
	"logSaver/pkg/store/mongostore"
	_ "logSaver/pkg/store/oraclestore"
)

type LogHandler struct {
	DB *mongostore.DB
}

func (lh *LogHandler) CreateLog(context *gin.Context) {
	logData := model.Log{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	err := lh.DB.LogRepository.Insert(logData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved to Mongo"})
}

func (lh *LogHandler) UpdateLog(context *gin.Context) {
	newLogData := model.Log{}
	if err := context.BindJSON(&newLogData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	existLogData, err := lh.DB.LogRepository.Get(&newLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	existLogData.UpdateExistLog(&newLogData)

	err = lh.DB.LogRepository.Update(existLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log updated"})
}
