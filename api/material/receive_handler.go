package material

import "github.com/gin-gonic/gin"
import (
	"androidappServer/pkg/structs"
)

type wirteReceiveBody struct {
	structs.Material
	ReceiverID int `json:"receiver"`

}

func WriteReceiveHandler(ctx *gin.Context){
	var reqBody wirteReceiveBody
	ctx.BindJSON(&reqBody)
}