package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"user-api/handler/user_addr"
	"user-api/handler/user_info"
	"user-api/handler/user_level"
)

//注册路由
func RegiserRouter(client *grpc.ClientConn, router *gin.Engine) {
	userRout := router.Group("/user")
	{
		apiService := user_info.Init(client)
		userRout.POST("/login", apiService.Login)
		userRout.POST("/out", apiService.Logout)
		userRout.POST("/register", apiService.Register)
		userRout.POST("/code", apiService.GetCode)
		userRout.POST("/list", apiService.GetUserInfoList)
		userRout.GET("/info", apiService.GetUserInfo)
		userAddrRout := userRout.Group("/addr")
		{
			userAddrService := user_addr.Init(client)
			userAddrRout.POST("/add", userAddrService.Add)
		}
		userLevelRout := userRout.Group("/level")
		{
			userLevelService := user_level.Init(client)
			userLevelRout.GET("/list", userLevelService.GetUserLevels)
		}
	}

}
