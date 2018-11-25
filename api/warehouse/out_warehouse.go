package warehouse

import (
	"androidappServer/internal/response"
	"androidappServer/internal/warehouse"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
