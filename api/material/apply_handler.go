package material

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
)

//填写申请表
func WriteApplyTableHandler(ctx *gin.Context) {
	var table material.ApplyTableBody
	var resBody response.ResBody
	defer ctx.JSON(http.StatusOK, &resBody)
	err := ctx.BindJSON(&table)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	// check param
	if !material.CheckApplyTableParam(&table) {
		logger.Info("check request parameter error")
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err = material.CreateApplyTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = fmt.Sprintf("store apply table to database error, %s", err.Error())
		logger.Error("store apply table to database error, error message: ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

// 申请单列表
func ApplyListhandler(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	var resBody response.ResBody
	defer ctx.JSON(http.StatusOK, &resBody)
	result, err := material.QueryApplyList(userID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query list failed"
		logger.Error("query list failed, error message: ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}

// 申请表详细信息
func ApplyDetailHandler(ctx *gin.Context) {
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
	result, err := material.QueryApplyDetail(tableID)
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
func ApplyVerifyHandler(ctx *gin.Context) {
	var body material.VerifyBody
	var resBody response.ResBody
	err := ctx.BindJSON(&body)
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	defer ctx.JSON(http.StatusOK, &resBody)
	if body.TableID == 0 || body.Status == "" || body.Verifier == 0{
		resBody.Status = status.StatusFailed
		resBody.Msg = "parameter check failed"
		return
	}
	err = material.ModifyApplyStatus(&body)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = fmt.Sprintf("verify failed, %s", err.Error())
		logger.Error("verify apply table failed, error message: ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}