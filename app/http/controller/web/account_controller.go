package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/middleware/request_context"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/model"
	"goskeleton/app/utils/logger"
	"goskeleton/app/utils/response"
)

type AccountController struct{}

func (u *AccountController) Login(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	loginReq, _ := req.(*common.LoginReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	account, err := model.CreateAccountFactory("").QueryByPhone(loginReq.Phone)
	if err != nil {
		zLogger.Error(err, "根据手机号查询账号失败")
		response.Success(context, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, "")
		return
	}

	if account.Id == 0 || account.Password != loginReq.Password {
		zLogger.Info(err, "登陆账号不存在或密码错误, account=[%v]", account)
		response.Success(context, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, "")
		return
	}

	zLogger.Info(nil, "登陆成功, account=[%v]", account)
	response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
}

func (u *AccountController) Register(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	registerReq, _ := req.(*common.RegisterReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	account, err := model.CreateAccountFactory("").QueryByPhone(registerReq.Phone)
	if err != nil {
		zLogger.Error(err, "根据手机号查询账号失败")
		response.Success(context, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
		return
	}

	if account.Id == 0 {
		zLogger.Info(err, "账号已注册, account=[%v]", account)
		response.Success(context, consts.CurdRegisterFailCode, "账号已注册", "")
		return
	}

	success, err := model.CreateAccountFactory("").
		StoreAccount(registerReq.Phone, registerReq.Email, registerReq.Password, registerReq.Code)
	if !success || err != nil {
		zLogger.Error(err, "存储账号信息失败")
		response.Success(context, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
		return
	}

	zLogger.Info(nil, "账号注册成功")
	response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
}
