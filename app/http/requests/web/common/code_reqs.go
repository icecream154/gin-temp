package common

type SendCodeReq struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
}

type ValidateCodeReq struct {
	Phone string `json:"phone" form:"phone" binding:"required"` //  手机号
	Code  string `json:"code" form:"code" binding:"required"`   // 验证码
}
