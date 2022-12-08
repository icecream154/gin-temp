package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/request_context"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/model"
	"goskeleton/app/utils/logger"
	"goskeleton/app/utils/response"
)

type SysOpinionController struct{}

func (u *SysOpinionController) SubmitOpinion(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	submitOpinionReq, _ := req.(*common.SubmitOpinionReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	accClaims := authorization.GetAccClaims(context)
	if !accClaims.IsAccClaimsValid() {
		response.ErrorTokenAuthFail(context)
		return
	}

	content := submitOpinionReq.Content
	contact := submitOpinionReq.Contact
	image := submitOpinionReq.Image

	success, err := model.CreateOpinionFactory("").StoreOpinion(accClaims.Account, content, image, contact)
	if !success || err != nil {
		zLogger.Error(err, "意见信息存储失败")
		response.Success(context, consts.SubmitOpinionFailCode, consts.SubmitOpinionFailMsg, "")
		return
	}

	zLogger.Info(nil, "意见信息提交成功")
	response.Success(context, consts.CurdStatusOkCode, consts.SubmitOpinionSuccessMsg, "")
}

func (o *SysOpinionController) QueryOpinion(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	qReq, _ := req.(*common.QueryOpinionReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(nil, requestContext)

	accClaims := authorization.GetAccClaims(context)
	if !accClaims.IsAccClaimsValid() {
		response.ErrorTokenAuthFail(context)
		return
	}

	total, list, err := model.CreateOpinionFactory("").QueryOpinion(qReq.Id,
		qReq.Account, qReq.Dealt, qReq.PageNum, qReq.PageSize)
	if err != nil {
		zLogger.Error(err, "查询意见失败")
		response.Success(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
		return
	}

	zLogger.Info(nil, "意见列表查询成功，total=%d，list=%v", total, list)
	response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, gin.H{
		"total": total,
		"list":  list,
	})
}
