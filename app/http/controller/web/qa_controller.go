package web

import (
	ctx "context"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/middleware/request_context"
	"goskeleton/app/http/requests/web/common"
	"goskeleton/app/utils/logger"
	"goskeleton/app/utils/response"
	"time"
)

type QAController struct{}

func (u *QAController) HandleInput(context *gin.Context) {
	req, _ := context.Get(consts.RequestKey)
	submitInputReq, _ := req.(*common.SubmitInputReq)

	requestContext := request_context.GetRequestContext(context)
	zLogger := logger.GetLogger(&req, requestContext)

	apiCallStart := time.Now().Unix()
	client := openai.NewClient("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")

	openAiChatMsgs := make([]openai.ChatCompletionMessage, len(submitInputReq.HistoryMsg)+1)
	for i := 0; i < len(submitInputReq.HistoryMsg); i++ {
		openAiChatMsgs[i] = openai.ChatCompletionMessage{
			Role:    submitInputReq.HistoryMsg[i].Role,
			Content: submitInputReq.HistoryMsg[i].Content,
		}
	}
	openAiChatMsgs[len(submitInputReq.HistoryMsg)] = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: submitInputReq.Content,
	}

	resp, err := client.CreateChatCompletion(
		ctx.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: openAiChatMsgs,
		},
	)
	apiCallEnd := time.Now().Unix()

	if err != nil {
		zLogger.Error(err, "open api call error")
	} else {
		zLogger.Info(nil, "open api call ok, resp=[%v]", resp)
	}

	if len(resp.Choices) > 0 {
		message := resp.Choices[0].Message.Content
		zLogger.Info(nil, "open api resp ok, msg=[%s]", message)
		response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, gin.H{
			"message":  message,
			"api_cost": apiCallEnd - apiCallStart,
		})
	} else {
		zLogger.Info(nil, "open api resp no content")
		response.Success(context, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, gin.H{
			"message": "",
		})
	}
}
