package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"phoenix/api"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", api.HealthHandler)   // 健康检查接口
	rg.GET("/info", api.InfoHandler)       // 开发者信息接口
	rg.GET("/version", api.VersionHandler) // 版本信息接口
	rg.POST("/login", auth.LoginHandler)   // 用户登录接口
	return rg
}

// 登录路由组
func LoginRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/logout", auth.LogoutHandler) // 用户注销接口
	return rg
}
