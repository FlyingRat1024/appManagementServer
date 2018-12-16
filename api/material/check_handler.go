package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 质检
func WriteCheckTableHandler(ctx *gin.Context) {
	var table material.CheckTableBody
	var resBody response.ResBody
	err := ctx.BindJSON(&table)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	defer ctx.JSON(http.StatusOK, &resBody)
	// check param
	if !material.CheckCheckTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err = material.CreateCheckTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store check table to database error"
		logger.Error("store check table to database error", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}