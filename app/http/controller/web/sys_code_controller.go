package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/request_context"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/model"
	"goskeleton/app/utils/logger"
	"goskeleton/app/utils/response"
	"math/rand"
	"strconv"
	"time"
)

// 有效时间两小时
const ExpireTime int64 = 2 * 60 * 60

const controllerName = "SysCodeController"

type SysCodeController struct{}

// 发送验证码
func (u *SysCodeController) SendCode(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	sendCodeReq, _ := req.(*common.SendCodeReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	beforeCode, err := model.CreateCodeFactory("").QueryUncheckedByPhone(sendCodeReq.Phone)
	if err != nil {
		zLogger.Error(err, "未校验的验证码查询失败")
		response.Success(context, consts.SendCodeFailCode, consts.SendCodeFailMsg, "")
		return
	}

	if beforeCode.Id != 0 {
		beforeIssueTime, _ := time.Parse(variable.DateFormatWithZone, beforeCode.IssueTime)
		left := 30 - (time.Now().Unix() - beforeIssueTime.Unix() + 8*3600)
		if time.Now().Unix()-beforeIssueTime.Unix()+8*3600 < 30 {
			zLogger.Warn(nil, "请求验证码过于频繁")
			response.Success(context, consts.SendCodeTooFreqCode, consts.SendCodeTooFreqMsg+"("+strconv.FormatInt(left, 10)+")", "")
			return
		}
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%04v", rnd.Int31n(10000))

	updateSuccess, updateErr := model.CreateCodeFactory("").StoreCode(sendCodeReq.Phone, code)
	if !updateSuccess || updateErr != nil {
		zLogger.Error(updateErr, "验证码存储失败，code=%s", code)
		response.Success(context, consts.SendCodeFailCode, consts.SendCodeFailMsg, updateErr)
		return
	}

	//err = sys_text_service.SendCode(sendCodeReq.Phone, code)
	//if err != nil {
	//	zLogger.Error(err, "阿里云验证码发送失败")
	//	response.Success(context, consts.SendCodeFailCode, consts.SendCodeFailMsg, err)
	//	return
	//}

	zLogger.Info(nil, "验证码发送成功，code=%s", code)
	response.Success(context, consts.CurdStatusOkCode, consts.SendCodeSuccessMsg, "")
}

// 校验验证码
func (u *SysCodeController) ValidateCode(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	validateCodeReq, _ := req.(*common.ValidateCodeReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	phone := validateCodeReq.Phone
	code := validateCodeReq.Code

	codeModel, err := model.CreateCodeFactory("").QueryUncheckedByPhoneAndCode(phone, code)
	if err != nil {
		zLogger.Error(err, "未校验的验证码查询失败")
		response.Success(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
		return
	}

	// 验证码无效
	if codeModel.Id == 0 {
		zLogger.Warn(nil, "用户验证码无效")
		response.Success(context, consts.PhoneCodeInvalidCode, consts.PhoneCodeInvalidMsg, "")
		return
	}

	// 验证码有效但已过期
	issueTime, _ := time.Parse(variable.DateFormatWithZone, codeModel.IssueTime)
	if time.Now().Unix()-issueTime.Unix() > ExpireTime {
		zLogger.Warn(nil, "用户验证码过期，codeModel=%v", codeModel)
		response.Success(context, consts.PhoneCodeExpiredCode, consts.PhoneCodeExpiredMsg, "")
		return
	}

	zLogger.Info(nil, "用户验证码校验正确，codeModel=%v", codeModel)
	response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
}
