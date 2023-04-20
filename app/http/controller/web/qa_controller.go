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
	resp, err := client.CreateChatCompletion(
		ctx.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: submitInputReq.Content,
				},
			},
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
