package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 用户路由组
func UserRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rgs := rg.Group("/user")
	rgs.Use(auth.MiddlewareFunc())                     // 登录中间件
	rgs.GET("/list", v1.UserListHandler)               // 获取用户列表接口
	rgs.GET("/info", v1.CurrentUserInfoHandler)        // 获取当前用户信息接口
	rgs.GET("/:jobId/info", v1.UserInfoByJobIdHandler) // 获取用户信息接口
	return rgs
}
