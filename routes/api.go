package routes

import (
	"github.com/gin-gonic/gin"
	"common_gin/app/controller"
	"common_gin/app/middleware"
)

func InitApiRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	//记录日志
	router.Use(middleware.RequestLog())
	//用户
	user := router.Group("/")
	{
		controller.UserRegister(user)
	}

	return router
}
