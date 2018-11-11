package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import (
	"androidappServer/internal"
)

type requestBody struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type responseBody struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Param  interface{} `json:"param"`
}

func LoginHandler(ctx *gin.Context) {
	//todo
	resBody := responseBody{
		Status: 0,
		Msg:    "success",
		Param:  "",
	}
	var reqBody requestBody
	ctx.BindJSON(&reqBody)
	flag, err := internal.Login(reqBody.ID, reqBody.Password)
	if err != nil {
		resBody.Msg = fmt.Sprintf("error: %s", err.Error())
		ctx.JSON(200, resBody)
		return
	}
	if flag {
		resBody.Status = 1
		resBody.Msg = "login success"
		ctx.JSON(200, resBody)
		return
	}
	resBody.Msg = "login failed"
	resBody.Param = reqBody
	ctx.JSON(200, resBody)
	return
}
