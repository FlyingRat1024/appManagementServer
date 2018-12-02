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
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "get material list failed"
		logger.Error("get material list failed, ", err.Error())
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "get material list success"
	resBody.Param = jsonStr
}

func CreateMaterialHandler(ctx *gin.Context) {
	resBody := response.ResBody{}
	var reqBody material.Material
	ctx.BindJSON(&reqBody)
	defer ctx.JSON(http.StatusAccepted, &resBody)
	if !material.CheckMaterialParam(&reqBody) {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check param error, please check your json string"
		return
	}
	err := material.CreateMaterial(&reqBody)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "execute database success"
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
}
