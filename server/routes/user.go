package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 用户路由组
func UserRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", v1.UserListHandler) // 用户列表接口
	return rg
}
