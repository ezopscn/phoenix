package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 菜单路由组
func MenuRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rgs := rg.Group("/menu")
	rgs.Use(auth.MiddlewareFunc())                            // 登录中间件
	rgs.GET("/list", v1.CurrentUserMenuListHandler)           // 当前用户角色的菜单列表接口
	rgs.GET("/role/:roleId/list", v1.MenuListByRoleIdHandler) // 指定角色的菜单列表接口
	return rgs
}
