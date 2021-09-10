package routers

import (
	"github.com/gin-gonic/gin"

	"im/internal/api/handlers"
	"im/internal/api/middlewares"
	"im/pkg/proto"
)

// todo 1 为 API 提供 SDK 支持

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(middlewares.CORS())
	initUserRouter(router)
	initChatRouter(router)
	router.NoRoute(func(c *gin.Context) {
		handlers.ResponseWithCode(c, proto.Code404, handlers.NewResponse("request uri not found", nil))
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

func initChatRouter(r *gin.Engine) {
	r.POST("/peer/chat", handlers.PeerChat)
	r.POST("/group/chat", handlers.GroupChat)
	r.POST("/group/count", handlers.GroupCount)
	r.POST("/group/info", handlers.GroupInfo)
	r.Use(middlewares.SessionCheck())
}
