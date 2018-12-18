package message

import (
	"androidappServer/internal/message"
	"androidappServer/internal/response"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MaterialWarningHandler(ctx *gin.Context) {
	resBody := response.ResBody{}
	defer ctx.JSON(http.StatusOK, &resBody)
	result, err := message.GetWarning()
	if err != nil {
		resBody.Status = 0
		resBody.Msg = "query message error"
		logger.Error("query message error, ", err.Error())
	}
	resBody.Status = 1
	resBody.Msg = "query message success"
	resBody.Param = result
}
