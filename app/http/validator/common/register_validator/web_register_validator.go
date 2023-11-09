package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/http/validator/web/common"
)

func WebRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()
	containers.Set("Pin", common.PinValidator{})
}
