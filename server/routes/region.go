package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "phoenix/api/v1"
)

// 地区路由组
func RegionRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/province/list", v1.ProvinceListHandler)                          // 省份列表接口
	rg.GET("/province/:provinceId/city/list", v1.CityListByProvinceIdHandler) // 地级市列表接口
	rg.GET("/city/:cityId/area/list", v1.AreaListByCityIdHandler)             // 行政区列表接口
	rg.GET("/area/:areaId/street/list", v1.StreetListByAreaIdHandler)         // 街道列表接口
	return rg
}
