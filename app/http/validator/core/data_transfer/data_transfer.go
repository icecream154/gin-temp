package data_transfer

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"time"
)

// 将验证器成员(字段)绑定到数据传输到上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/

func DataAddContext(request interface{}, context *gin.Context) *gin.Context {
	context.Set(consts.RequestKey, request)
	curDateTime := time.Now().Format(variable.DateFormat)
	context.Set(consts.ValidatorPrefix+"created_at", curDateTime)
	context.Set(consts.ValidatorPrefix+"updated_at", curDateTime)
	context.Set(consts.ValidatorPrefix+"deleted_at", curDateTime)
	return context
}
