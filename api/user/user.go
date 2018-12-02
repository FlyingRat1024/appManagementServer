package user

import (
	"androidappServer/internal/response"
	"androidappServer/internal/user"
	"androidappServer/pkg/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserInfo(ctx *gin.Context) {
	var resBody response.ResBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	userIdStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	result, err := user.QueryUserInfo(userID)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query database error"
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	return
}

func ModifyPassword(ctx *gin.Context) {
	var resBody response.ResBody
	var reqBody user.ModifyPasswordBody
	defer ctx.JSON(http.StatusAccepted, &resBody)
	err := ctx.BindJSON(&reqBody)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	if reqBody.Password == "" || reqBody.UserID == 0 {
		resBody.Status = status.StatusFailed
		resBody.Msg = "check request parameter error"
		return
	}
	err = user.ModifyPassword(&reqBody)
	if err != nil {
		resBody.Status = status.StatusFailed
		resBody.Msg = "query database error"
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	return
}
