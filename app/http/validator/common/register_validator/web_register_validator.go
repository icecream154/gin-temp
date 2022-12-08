package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/http/validator/web/common"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	// 发送与校验验证码
	containers.Set("SysSendCode", common.SendCodeValidator{})
	containers.Set("SysValidateCode", common.ValidateCodeValidator{})

	// 意见提交与查询
	containers.Set("SysSubmitOpinion", common.SubmitOpinionValidator{})
	containers.Set("SysQueryOpinion", common.QueryOpinionValidator{})
}
