package controller

import (
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
		Status: 1,
		Msg:"success",
		Param:"",
	}
	param, err := internal.Login()
	if err != nil{
		resBody.Param = param
	}
	ctx.JSON(200,resBody )
}
