package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

//填写申请表
func WriteApplyTableHandler(ctx *gin.Context) {
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
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store apply table to database error"
		logger.Error("store apply table to database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

// 申请单列表
func ApplyListhandler(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	result, err := material.QueryApplyList()
	if err != nil{
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

// 申请表详细信息
func ApplyDetailHandler(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	tableID := ctx.GetInt("table_id")
	if tableID == 0{
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	result, err := material.QueryApplyDetail(tableID)
	if err != nil{
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
