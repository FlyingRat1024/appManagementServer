package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func WriteReceiveHandler(ctx *gin.Context) {
	var table material.RecieveTableBody
	var resBody response.ResBody
	err := ctx.BindJSON(&table)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	defer ctx.JSON(http.StatusOK, &resBody)
	// check param
	if !material.CheckReceiveTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err = material.CreateReceiveTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = fmt.Sprintf("store receive table to database error, %s", err.Error())
		logger.Error("store receive table to database error", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

// 领料单列表
func ReceiveListhandler(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusOK, &resBody)
	userID := ctx.Query("user_id")
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
	defer ctx.JSON(http.StatusOK, &resBody)
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
		resBody.Msg = fmt.Sprintf("query detail failed, %s", err.Error())
		logger.Error("query apply detail failed, error message: ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}

// 经理审核
func ReceiveVerifyHandler(ctx *gin.Context) {
	var body material.VerifyBody
	var resBody response.ResBody
	err := ctx.BindJSON(&body)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	defer ctx.JSON(http.StatusOK, &resBody)
	if body.TableID == 0 || body.Status == ""{
		resBody.Status = status.StatusFailed
		resBody.Msg = "parameter check failed"
		return
	}
	err = material.ModifyReceiveStatus(&body)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = fmt.Sprintf("verify failed, %s", err.Error())
		logger.Error("verify receive table failed, error message: ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}
