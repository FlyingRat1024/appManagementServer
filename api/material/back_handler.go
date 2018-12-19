package material

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
)

func WriteBackTableHandler(ctx *gin.Context) {
	var table material.BackTableBody
	var resBody response.ResBody
	err := ctx.BindJSON(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	defer ctx.JSON(http.StatusOK, &resBody)
	// check param
	if !material.CheckBackTableParam(&table) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	//store db
	err = material.CreateBackTable(&table)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = fmt.Sprintf("store back table to database error, %s", err.Error())
		logger.Error("store back table to database error", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}
