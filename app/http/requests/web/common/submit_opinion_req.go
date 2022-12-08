package common

type SubmitOpinionReq struct {
	Content string `form:"content" json:"content" binding:"required"`
	Contact string `form:"contact" json:"contact"`
	Image   string `form:"image" json:"image"`
}

type QueryOpinionReq struct {
	Id       int64  `form:"id" json:"id"`
	Account  string `form:"account" json:"account"`
	Dealt    int8   `form:"dealt" json:"dealt"`
	PageNum  int    `form:"page_num" json:"page_num" binding:"required"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required"`
}
