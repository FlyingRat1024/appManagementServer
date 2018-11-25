package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteReceiveHandler(ctx *gin.Context){
	var table material.RecieveTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	// check param
	if !material.CheckReceiveTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err := material.CreateReceiveTable(&table)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "store receive table to database error"
		logger.Error("store receive table to database error", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}
