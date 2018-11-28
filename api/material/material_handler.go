package material

import (
	"androidappServer/internal/material"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MaterialListHandler(ctx *gin.Context) {
	resBody := response.ResBody{}
	defer ctx.JSON(http.StatusAccepted, &resBody)
	jsonStr, err := material.GetMaterialList()
	if err != nil{
		resBody.Status = status.StatusFailed
		resBody.Msg = "get material list failed"
		logger.Error("get material list failed, ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "get material list success"
	resBody.Param = jsonStr
}