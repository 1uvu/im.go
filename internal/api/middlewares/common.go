package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"im/internal/api/handlers"
	"im/internal/pkg/rpc"
	"im/pkg/config"
	"im/pkg/proto"
)

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.APISessionCheckRequest
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			c.Abort()
			handlers.ResponseWithCode(c, proto.CodeSessionError, proto.APIResponse{})
			return
		}

		reply := new(proto.LogicAuthCheckReply)

		ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
			"AuthCheck",
			&proto.LogicAuthCheckArg{
				AuthToken: req.AuthToken,
			},
			reply,
			func(reply proto.IRPCReply) bool {
				_reply := reply.(*proto.LogicAuthCheckReply)
				return _reply.Code != proto.CodeFailedReply && _reply.UserID > 0 && _reply.UserName != ""
			},
		)

		if !ok {
			c.Abort()
			handlers.ResponseWithCode(c, proto.CodeSessionError, proto.APIResponse{
				Msg:  reply.GetErrMsg(),
				Data: nil,
			})
			return
		}

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.GetConfig().API.CORSFlag {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}

		method := c.Request.Method
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}

		c.Next()
	}
}
