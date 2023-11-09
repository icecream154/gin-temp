package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/utils/response"
)

type PinController struct{}

func (u *PinController) Pin(context *gin.Context) {
	response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "pin success")
}
