package routers

import (
	"github.com/gin-gonic/gin"

	"im/internal/api/handlers"
	"im/internal/api/middlewares"
	"im/pkg/config"
	"im/pkg/proto"
)

// opt 为 API 提供 SDK 支持

var router *gin.Engine

func init() {
	router = gin.Default()

	if config.GetConfig().API.CORSFlag {
		router.Use(middlewares.CORS())
	}

	initUserRouter(router)
	initPushRouter(router)
	router.NoRoute(func(c *gin.Context) {
		handlers.ResponseWithCode(c, proto.Code404, proto.APIResponse{
			Msg:  "request uri not found",
			Data: nil,
		})
	})
}

func GetRouter() *gin.Engine {
	return router
}

func initUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/signin", handlers.Signin)
	userGroup.POST("/signup", handlers.Signup)
	userGroup.POST("/signout", handlers.Signout)
	userGroup.POST("/authcheck", handlers.AuthCheck)
	userGroup.Use(middlewares.SessionCheck())
}

func initPushRouter(r *gin.Engine) {
	r.POST("/peer/push", handlers.PeerPush)
	r.POST("/group/push", handlers.GroupPush)
	r.POST("/group/count", handlers.GroupCount)
	r.POST("/group/info", handlers.GroupInfo)
	r.Use(middlewares.SessionCheck())
}
