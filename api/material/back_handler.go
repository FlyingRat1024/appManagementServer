package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteBackTableHandler(ctx *gin.Context) {
	var table material.BackTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	// check param
	if !material.CheckBackTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err := material.CreateBackTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store back table to database error"
		logger.Error("store back table to database error", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}
