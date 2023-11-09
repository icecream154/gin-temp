package common

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
)

type PinValidator struct {
}

func (l PinValidator) CheckParams(context *gin.Context) {
	extraAddBindDataContext := data_transfer.DataAddContext(nil, context)
	(&web.PinController{}).Pin(extraAddBindDataContext)
}
