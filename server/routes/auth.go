package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 认证路由组
func AuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.Use(auth.MiddlewareFunc()) // 登录中间件

	// 其它
	rg.GET("/logout", auth.LogoutHandler) // 用户注销接口

	// 地区路由
	rg.GET("/region/province/list", v1.ProvinceListHandler)                          // 获取省份列表接口
	rg.GET("/region/province/:provinceId/city/list", v1.CityListByProvinceIdHandler) // 获取地级市列表接口
	rg.GET("/region/city/:cityId/area/list", v1.AreaListByCityIdHandler)             // 获取行政区列表接口
	rg.GET("/region/area/:areaId/street/list", v1.StreetListByAreaIdHandler)         // 获取街道列表接口

	// 用户路由
	rg.GET("/user/list", v1.GetUserListHandler)               // 获取用户列表接口
	rg.GET("/user/info", v1.GetCurrentUserInfoHandler)        // 获取当前用户的用户信息接口
	rg.GET("/user/:jobId/info", v1.GetSpecifyUserInfoHandler) // 获取指定用户的用户信息接口

	// 角色路由
	rg.GET("/role/list", v1.GetRoleListHandler)                           // 获取角色列表接口
	rg.GET("/role/:roleKeyword/info", v1.GetRoleInfoByRoleKeywordHandler) // 获取指定角色的信息接口

	// 菜单路由
	rg.GET("/menu/list", v1.GetMenuListHandler)                                         // 获取所有菜单列表接口
	rg.GET("/menu/list/for/role/:roleKeyword", v1.GetSpecifyRoleKeywordMenuListHandler) // 获取指定角色的菜单列表接口
	rg.GET("/menu/:menuId/info", v1.GetMenuInfoHandler)                                 // 获取指定菜单的详细信息接口

	return rg
}
