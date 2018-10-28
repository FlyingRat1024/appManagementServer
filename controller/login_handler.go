package controller

import "github.com/gin-gonic/gin"

func LoginHandler(ctx *gin.Context) {
	//todo
	ctx.JSON(200,gin.H{
		"msg":"success",
	})
}
