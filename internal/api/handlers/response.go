package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"im/pkg/common"
	"im/pkg/proto"
)

func Success(c *gin.Context, msg string, data common.Any) {
	ResponseWithCode(c, proto.CodeSuccessReply, proto.APIResponse{
		Msg:  msg,
		Data: data,
	})
}

func Failed(c *gin.Context, msg string, data common.Any) {
	ResponseWithCode(c, proto.CodeFailedReply, proto.APIResponse{
		Msg:  msg,
		Data: data,
	})
}

func ResponseWithCode(c *gin.Context, code int, resp proto.APIResponse) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": proto.CodeText(code),
		"msg":  resp.Msg,
		"data": resp.Data,
	})
}
