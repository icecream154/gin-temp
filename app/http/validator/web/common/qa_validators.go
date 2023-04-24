package common

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type SubmitInputValidator struct {
}

func (l SubmitInputValidator) CheckParams(context *gin.Context) {
	var submitInputReq common.SubmitInputReq
	if err := context.ShouldBind(&submitInputReq); err != nil {
		errs := gin.H{
			"tips": "参数校验失败",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	if submitInputReq.HistoryMsg == nil {
		submitInputReq.HistoryMsg = make([]common.ChatMsg, 0)
	}

	extraAddBindDataContext := data_transfer.DataAddContext(&submitInputReq, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "参数绑定失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.QAController{}).HandleInput(extraAddBindDataContext)
	}
}
