package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 角色路由组
func RoleRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rgs := rg.Group("/role")
	rgs.Use(auth.MiddlewareFunc())                        // 登录中间件
	rgs.GET("/list", v1.RoleListHandler)                  // 获取角色列表接口
	rgs.GET("/role/:roleId/info", v1.RoleInfoByIdHandler) // 获取指定角色的信息接口
	return rgs
}
