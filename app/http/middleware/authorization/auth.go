package authorization

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/my_jwt"
	token2 "goskeleton/app/service/token"
)

type AuthHeaderParams struct {
	Authorization string `header:"Authorization"`
}

func GetAccClaims(context *gin.Context) my_jwt.AccClaims {
	key := variable.ConfigYml.GetString("Token.BindContextKeyName")
	// token验证通过，同时绑定在请求上下文
	val, exist := context.Get(key)
	if !exist {
		return my_jwt.AccClaims{}
	}
	accClaims, _ := val.(my_jwt.AccClaims)
	return accClaims
}

func GetRawToken(context *gin.Context) string {
	val, exist := context.Get("Token.Raw")
	if !exist {
		return ""
	}
	rawToken, _ := val.(string)
	return rawToken
}

// 检查token权限
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeaderParams := AuthHeaderParams{}
		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&authHeaderParams); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsValidatorBindParamsFail, zap.Error(err))
			context.Abort()
			return
		}

		if len(authHeaderParams.Authorization) >= 20 {
			token := authHeaderParams.Authorization
			tokenIsEffective := token2.CreateAccTokenFactory().IsEffective(token)
			if tokenIsEffective {
				if accClaims, err := token2.CreateAccTokenFactory().ParseToken(token); err == nil {
					key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					context.Set(key, accClaims)
					context.Set("Token.Raw", authHeaderParams.Authorization)
				}
			}
		}
		context.Next()
	}
}
