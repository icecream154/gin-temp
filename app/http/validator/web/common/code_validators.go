package common

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type LoginValidator struct {
}

func (l LoginValidator) CheckParams(context *gin.Context) {
	var loginReq common.LoginReq
	if err := context.ShouldBind(&loginReq); err != nil {
		errs := gin.H{
			"tips": "参数校验失败",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&loginReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.AccountController{}).Login(extraAddBindDataContext)
	}
}

type RegisterValidator struct {
}

func (l RegisterValidator) CheckParams(context *gin.Context) {
	var registerReq common.RegisterReq
	if err := context.ShouldBind(&registerReq); err != nil {
		errs := gin.H{
			"tips": "手机号或验证码参数缺失",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&registerReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.AccountController{}).Register(extraAddBindDataContext)
	}
}
