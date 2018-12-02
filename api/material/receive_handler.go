package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func WriteReceiveHandler(ctx *gin.Context) {
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
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store receive table to database error"
		logger.Error("store receive table to database error", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

// 领料单列表
func ReceiveListhandler(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	userIdStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	result, err := material.QueryReceiveTableList(userID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query list failed"
		logger.Error("query list failed, error message: ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}

// 领料表详细信息
func ReceiveDetailHandler(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	tableIDStr := ctx.Query("table_id")
	if tableIDStr == "" {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	tableID, err := strconv.Atoi(tableIDStr)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	result, err := material.QueryReceiveDetail(tableID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query detail failed"
		logger.Error("query apply detail failed, error message: ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}
