package response

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/request_context"
	"goskeleton/app/utils/logger"
	"net/http"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	Context.Header("TraceId", request_context.GetRequestContext(Context).TraceId)
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// http返回成功
func Success(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, dataCode, msg, data)
}

// 请求失败
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

//token 权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	requestContext := request_context.GetRequestContext(c)
	zLogger := logger.GetLogger(nil, requestContext)
	accClaims := authorization.GetAccClaims(c)

	zLogger.Warn(nil, "Token Auth Fail For Path [%s]: Authorization=[%s] and AccClaims=[%v]",
		c.Request.RequestURI, c.Request.Header.Get("Authorization"), accClaims)
	ReturnJson(c, http.StatusOK, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// casbin 鉴权失败，返回 405 方法不允许访问
func ErrorCasbinAuthFail(c *gin.Context, msg interface{}) {
	ReturnJson(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, my_errors.ErrorsCasbinNoAuthorization, msg)
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusOK, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}
