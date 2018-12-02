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

func WriteOutWarehouseHandler(ctx *gin.Context) {
	var table warehouse.OutWarehouseTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	if !warehouse.CheckOutWarehouseTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err := warehouse.CreateOutWarehouseTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "store out warehouse table to database error"
		logger.Error("store out warehouse table to database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

func ConfirmOutWarehouseHandler(ctx *gin.Context) {
	var reqBody warehouse.ConfirmOutWarehouseBody
	var resBody response.ResBody
	ctx.BindJSON(&reqBody)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	if !warehouse.CheckConfirmOutWarehouseParam(&reqBody) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//update db
	err := warehouse.ConfirmOutWarehouseTable(&reqBody)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "modify database error"
		logger.Error("confirm out warehouse table, update  database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}

func OutWarehouseListHandler(ctx *gin.Context) {
	userIdStr := ctx.Query("user_id")
	var resBody response.ResBody
	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	result, err := warehouse.QueryOutWarehouseList(userID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query database error"
		ctx.JSON(http.StatusAccepted, &resBody)
		logger.Error("query out warehouse list from database error, error message ", err)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	ctx.JSON(http.StatusAccepted, &resBody)
	return
}

func OutWarehouseDetailHandler(ctx *gin.Context) {
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