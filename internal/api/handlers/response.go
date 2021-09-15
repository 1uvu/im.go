package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"im/pkg/common"
	"im/pkg/proto"
)

type Response struct {
	Msg  string
	Data common.Any
}

func NewResponse(msg string, data common.Any) Response {
	return Response{
		Msg:  msg,
		Data: data,
	}
}

func Success(c *gin.Context, resp Response) {
	ResponseWithCode(c, proto.CodeSuccessReply, resp)
}

func Failed(c *gin.Context, resp Response) {
	ResponseWithCode(c, proto.CodeFailedReply, resp)
}

func ResponseWithCode(c *gin.Context, code int, resp Response) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": proto.CodeText(code),
		"msg":  resp.Msg,
		"data": resp.Data,
	})
}
