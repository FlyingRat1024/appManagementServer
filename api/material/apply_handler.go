package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)


func WriteApplyTableHandler (ctx *gin.Context){
	var table material.ApplyTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	// check param
	if !material.CheckApplyTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err := material.CreateApplyTable(&table)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "store apply table to database error"
		logger.Error("store apply table to database error", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}