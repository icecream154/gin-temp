package common

type LoginReq struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `json:"password" form:"password"`
}

type RegisterReq struct {
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Code     string `json:"code" form:"code"` // 邀请码
}
