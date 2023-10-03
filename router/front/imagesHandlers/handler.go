package imagesHandlers

import (
	authGlobal "chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/router/base"
	"chatgpt-web-new-go/service/draw"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	gogpt "github.com/sashabaranov/go-openai"
)

func ImageList(c *gin.Context) {
	request := &ImagesRequest{}

	if err := c.Bind(request); err != nil {
		logs.Error("request bind error: %v", err)
		base.Fail(c, "参数异常:"+err.Error())
		return
	}

	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	drawRecords, count, err := draw.DrawList(c, u.ID, request.Type, request.Page, request.PageSize)
	if err != nil {
		logs.Error("draw list error: %v", err)
		base.Fail(c, "查询异常："+err.Error())
		return
	}

	base.Success(c, base.PageResponse{
		Count: count,
		Rows:  drawRecords,
	})
}

func ImagesGenerationsHandler(c *gin.Context) {
	request := &draw.Request{}

	if err := c.BindJSON(request); err != nil {
		logs.Error("draw bind error: %v", err)
		base.Fail(c, "参数异常"+err.Error())
		return
	}

	uFromCtx, found := c.Get(authGlobal.GinCtxKey)
	if !found {
		logs.Error("cannot get auth user key:%v", authGlobal.GinCtxKey)
		base.Fail(c, "cannot get auth user key:"+authGlobal.GinCtxKey)
		return
	}

	u := uFromCtx.(*model.User)

	response, err := draw.Process(c, request, u.ID)
	if err != nil {
		logs.Error("draw process error: %v", err)
		base.Fail(c, "绘画服务异常："+err.Error())
		return
	}
	resp := func() []draw.Response {
		rs := make([]draw.Response, 0)
		for _, datum := range response.Data {
			rs = append(rs, draw.Response{
				ID:         grand.S(22),
				UserId:     gconv.String(u.ID),
				CreateTime: gtime.Now().String(),
				TakeTime:   response.Created,
				Status:     0,
				Model:      "",
				Prompt:     request.Prompt,
				Images:     []string{datum.URL},
				Size:       gogpt.CreateImageSize256x256,
			})
		}

		return rs
	}()

	base.Success(c, resp)
}
