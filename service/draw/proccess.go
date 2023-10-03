package draw

import (
	"chatgpt-web-new-go/common/aiClient"
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	gogpt "github.com/sashabaranov/go-openai"
)

func Process(ctx *gin.Context, r *Request, uid int64) (imageResp *gogpt.ImageResponse, err error) {
	imageResp = &gogpt.ImageResponse{}
	request := gogpt.ImageRequest{
		Prompt:         r.Prompt,
		N:              gconv.Int(r.Quantity),
		Size:           gogpt.CreateImageSize256x256,
		ResponseFormat: gogpt.CreateImageResponseFormatURL,
		User:           gconv.String(uid),
	}
	gptClient := aiClient.GetGptClient()
	if gptClient == nil {
		logs.Error("gptClient is nil!")
		return nil, bizError.AiKeyNoneUsefullError
	}
	*imageResp, err = gptClient.OpenAIClient.CreateImage(ctx, request)
	// insert message
	//goUtil.New(func() {
	//	addMessage(ctx, r, uid)
	//})
	//
	//// insert record
	//goUtil.New(func() {
	//	insertTurnOverRecord(ctx, r, uid)
	//})

	return
}

//func addMessage(ctx *gin.Context, r *Request, uid int64) {
//	msg := &model.Message{
//		UserID:           uid,
//		Content:          r.Prompt,
//		PersonaID:        types.InterfaceToInt64(r.PersonaId),
//		Role:             gogpt.ChatMessageRoleUser,
//		FrequencyPenalty: int32(r.Options.FrequencyPenalty),
//		MaxTokens:        int32(r.Options.MaxTokens),
//		Model:            r.Options.Model,
//		PresencePenalty:  int32(r.Options.PresencePenalty),
//		Temperature:      int32(r.Options.Temperature),
//		ParentMessageID:  r.ParentMessageId,
//	}
//
//	dm := dao.Q.Message
//	err := dm.WithContext(ctx).Create(msg)
//	if err != nil {
//		logs.Error("message create error: %v", err)
//	}
//	return
//}
