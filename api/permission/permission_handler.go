package permission

import (
	"androidappServer/internal/permission"
	"androidappServer/internal/response"
	"androidappServer/pkg/status"
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionHandler(ctx *gin.Context) {
	resBody := response.ResBody{
		Status: status.StatusFailed,
		Msg:    "",
		Param:  "",
	}
	role := ctx.Query("role")
	if role == "" {
		resBody.Msg = "login error, cant't get parameter `role` "
		ctx.JSON(http.StatusOK, resBody)
		return
	}
	result, err := permission.QueryPermission(role)
	if err != nil {
		logger.Error("get permission error, error message: ", err)
		resBody.Msg = fmt.Sprintf("get permission error: %s", err)
		ctx.JSON(http.StatusOK, resBody)
		return
	}
	resBody.Status = status.StatusSuccess
	resBody.Msg = "success"
	resBody.Param = result
	ctx.JSON(http.StatusOK, resBody)
	return
}

