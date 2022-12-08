package common

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type SendCodeValidator struct {
}

func (l SendCodeValidator) CheckParams(context *gin.Context) {
	var sendCodeReq common.SendCodeReq
	if err := context.ShouldBind(&sendCodeReq); err != nil {
		errs := gin.H{
			"tips": "参数校验失败",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&sendCodeReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.SysCodeController{}).SendCode(extraAddBindDataContext)
	}
}

type ValidateCodeValidator struct {
}

func (l ValidateCodeValidator) CheckParams(context *gin.Context) {
	var validateCodeReq common.ValidateCodeReq
	if err := context.ShouldBind(&validateCodeReq); err != nil {
		errs := gin.H{
			"tips": "手机号或验证码参数缺失",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&validateCodeReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.SysCodeController{}).ValidateCode(extraAddBindDataContext)
	}
}
