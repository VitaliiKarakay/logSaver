package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"logSaver/pkg/model"
	"logSaver/pkg/store/oraclestore"
)

type LogHandler struct {
	DB *oraclestore.DB
}

func (lh *LogHandler) CreateSMSLog(context *gin.Context) {
	logData := model.SMSLog{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	err := lh.DB.SmsRepository.InsertSMSLog(logData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}

func (lh *LogHandler) UpdateSMSLog(context *gin.Context) {
	newLogData := model.SMSLog{}
	if err := context.BindJSON(&newLogData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	existLogData, err := lh.DB.SmsRepository.GetSMSLog(&newLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	existLogData.UpdateExistLog(&newLogData)

	err = lh.DB.SmsRepository.UpdateSMSLog(existLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log updated"})
}

func (lh *LogHandler) CreateEmailLog(context *gin.Context) {
	logData := model.EmailLog{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	err := lh.DB.EmailRepository.InsertEmailLog(logData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}

func (lh *LogHandler) UpdateEmailLog(context *gin.Context) {
	newLogData := model.EmailLog{}
	if err := context.BindJSON(&newLogData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	existLogData, err := lh.DB.EmailRepository.GetEmailLog(&newLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	existLogData.UpdateExistLog(&newLogData)

	err = lh.DB.EmailRepository.UpdateEmailLog(existLogData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log updated"})
}
