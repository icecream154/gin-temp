package common

type ChatMsg struct {
	Role    string `form:"role" json:"role"`
	Content string `form:"content" json:"content"`
}

type SubmitInputReq struct {
	Phone          string    `form:"phone" json:"phone"`
	ConversationId string    `form:"conversation_id" json:"conversation_id"`
	Content        string    `form:"content" json:"content"`
	HistoryMsg     []ChatMsg `form:"history_msg" json:"history_msg"`
}
