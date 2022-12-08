package common

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type SubmitOpinionValidator struct {
}

func (l SubmitOpinionValidator) CheckParams(context *gin.Context) {
	var submitOpinionReq common.SubmitOpinionReq
	if err := context.ShouldBind(&submitOpinionReq); err != nil {
		errs := gin.H{
			"tips": "参数校验失败",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&submitOpinionReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.SysOpinionController{}).SubmitOpinion(extraAddBindDataContext)
	}
}

type QueryOpinionValidator struct {
}

func (l QueryOpinionValidator) CheckParams(context *gin.Context) {
	var queryOpinionReq common.QueryOpinionReq
	if err := context.ShouldBind(&queryOpinionReq); err != nil {
		errs := gin.H{
			"tips": "手机号或验证码参数缺失",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&queryOpinionReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.SysOpinionController{}).QueryOpinion(extraAddBindDataContext)
	}
}
