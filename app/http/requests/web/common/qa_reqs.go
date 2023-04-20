package common

type SubmitInputReq struct {
	Phone          string `form:"phone" json:"phone"`
	ConversationId string `json:"conversation_id" json:"conversation_id"`
	Content        string `form:"content" json:"content"`
}
