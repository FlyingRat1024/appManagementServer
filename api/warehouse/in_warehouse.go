package warehouse

import (
	"androidappServer/internal/response"
	"androidappServer/internal/warehouse"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteInWarehouseHandler(ctx *gin.Context) {
	var table warehouse.InWarehouseTableBody
	var resBody response.ResBody
	ctx.BindJSON(&table)
	defer ctx.JSON(http.StatusAccepted, &resBody)
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
