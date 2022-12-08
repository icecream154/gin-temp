package request_context

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
)

const (
	requestContextKey = "requestContextKey"
)

type RequestContext struct {
	TraceId       string `header:"TraceId"`
	Authorization string `header:"Authorization"`
}

func GetRequestContext(context *gin.Context) *RequestContext {
	val, exist := context.Get(requestContextKey)
	if !exist {
		fakeContext := RequestContext{}
		return &fakeContext
	}
	requestContext, _ := val.(RequestContext)
	return &requestContext
}

// 检查请求上下文
func CheckRequestContext() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestContext := RequestContext{}
		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&requestContext); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsValidatorBindParamsFail, zap.Error(err))
			context.Abort()
			return
		}
		if requestContext.TraceId == "" {
			u1, _ := uuid.NewV4()
			requestContext.TraceId = u1.String()
		}

		context.Set(requestContextKey, requestContext)
		context.Next()
	}
}
