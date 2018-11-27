package Login

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)
import (
	"androidappServer/internal/login"
	"androidappServer/pkg/status"
)

type requestBody struct {
	ID       string `json:"userID"`
	Password string `json:"password"`
}

type responseBody struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Param  interface{} `json:"param"`
}

func LoginHandler(ctx *gin.Context) {
	resBody := responseBody{
		Status: status.StatusFailed,
		Msg:    "",
		Param:  "",
	}
	var reqBody requestBody
	ctx.BindJSON(&reqBody)
	flag, err := login.Login(reqBody.ID, reqBody.Password)
	if err != nil {
		logger.Error("login error, error message: ", err)
		resBody.Msg = fmt.Sprintf("login error: %s", err)
		ctx.JSON(http.StatusOK, resBody)
		return
	}
	if flag {
		resBody.Status = status.StatusSuccess
		resBody.Msg = "login success"
		ctx.JSON(http.StatusOK, resBody)
		return
	}
	resBody.Msg = "login failed"
	resBody.Param = reqBody
	ctx.JSON(http.StatusOK, resBody)
	return
}
