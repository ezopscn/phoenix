package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 认证路由组
func AuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/logout", auth.LogoutHandler)                                            // 用户注销
	rg.GET("/region/province/list", v1.ProvinceListHandler)                          // 获取省份列表
	rg.GET("/region/province/:provinceId/city/list", v1.CityListByProvinceIdHandler) // 获取地级市列表
	rg.GET("/region/city/:cityId/area/list", v1.AreaListByCityIdHandler)             // 获取行政区列表
	rg.GET("/region/area/:areaId/street/list", v1.StreetListByAreaIdHandler)         // 获取街道列表
	rg.GET("/user/count", v1.GetUserCountHandler)                                    // 获取用户总数
	rg.GET("/user/info", v1.GetCurrentUserInfoHandler)                               // 获取当前用户的用户详情
	rg.GET("/role/info", v1.GetCurrentUserRoleInfoHandler)                           // 获取当前用户的角色详情
	rg.GET("/menu/tree", v1.GetCurrentUserMenuTreeHandler)                           // 获取当前用户的菜单树
	rg.GET("/department/info", v1.GetCurrentUserDepartmentInfo)                      // 获取当前用户的部门信息

	return rg
}

// 鉴权路由
func CasbinRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/user/list", v1.GetUserListHandler)                              // 获取用户列表
	rg.GET("/user/:jobId/info", v1.GetSpecifyUserInfoHandler)                // 获取指定用户的详情
	rg.GET("/role/list", v1.GetRoleListHandler)                              // 获取角色列表
	rg.GET("/role/:roleKeyword/menu/list", v1.GetSpecifyRoleMenuListHandler) // 获取指定角色的菜单列表
	rg.GET("/role/:roleKeyword/info", v1.GetSpecifyRoleInfoHandler)          // 获取指定角色的详情
	rg.GET("/menu/list", v1.GetMenuListHandler)                              // 获取菜单列表
	rg.GET("/menu/:menuId/info", v1.GetSpecifyMenuInfoHandler)               // 获取指定菜单的详情

	return rg
}
