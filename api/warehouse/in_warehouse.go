package warehouse

import (
	"androidappServer/internal/response"
	"androidappServer/internal/warehouse"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func WriteInWarehouseHandler(ctx *gin.Context) {
	var table warehouse.InWarehouseTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusOK, &resBody)
	if !warehouse.CheckInWarehouseTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err := warehouse.CreateInWarehouseTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store in warehouse table to database error"
		logger.Error("store in warehouse table to database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

// 入库单列表
func InWarehouseListHandler(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	var resBody response.ResBody
	result, err := warehouse.QueryInWarehouseList(userID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query database error"
		ctx.JSON(http.StatusOK, &resBody)
		logger.Error("query in warehouse list from database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	ctx.JSON(http.StatusOK, &resBody)
	return
}

//入库单详细信息
func InWarehouseDetailHandler(ctx *gin.Context) {
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
	result, err := warehouse.QueryInWarehouseDetail(tableID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query detail failed"
		logger.Error("query in warehouse detail failed, error message: ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}
