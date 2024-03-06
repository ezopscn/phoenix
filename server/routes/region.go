package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 地区路由组
func RegionRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rgs := rg.Group("/region")
	rgs.Use(auth.MiddlewareFunc())                                             // 登录中间件
	rgs.GET("/province/list", v1.ProvinceListHandler)                          // 获取省份列表接口
	rgs.GET("/province/:provinceId/city/list", v1.CityListByProvinceIdHandler) // 获取地级市列表接口
	rgs.GET("/city/:cityId/area/list", v1.AreaListByCityIdHandler)             // 获取行政区列表接口
	rgs.GET("/area/:areaId/street/list", v1.StreetListByAreaIdHandler)         // 获取街道列表接口
	return rgs
}
