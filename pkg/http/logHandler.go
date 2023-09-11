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

	err := lh.DB.LogRepository.Insert(logData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}

func (lh *LogHandler) UpdateLog(context *gin.Context) {
	logData := model.Log{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

}
