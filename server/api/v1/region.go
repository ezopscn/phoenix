package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
)

// 获取所有省份
func ProvinceListHandler(ctx *gin.Context) {
	var provinces []model.Province
	err := common.DB.Find(&provinces).Error
	if err != nil {
		response.FailedWithMessage("获取省份信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": provinces,
	})
}

// 根据省份获取地级市
func CityListByProvinceIdHandler(ctx *gin.Context) {
	// 参数处理
	sid := ctx.Param("provinceId")
	uid, err := utils.ConvertStringToUint(sid)
	if err != nil {
		response.FailedWithMessage(response.ParamErrorMessage)
		return
	}

	// 查询数据
	var cities []model.City
	err = common.DB.Where("province_id = ?", uid).Find(&cities).Error
	if err != nil {
		response.FailedWithMessage("获取地级市信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": cities,
	})
}

// 根据地级市获取区
func AreaListByCityIdHandler(ctx *gin.Context) {
	// 参数处理
	sid := ctx.Param("cityId")
	uid, err := utils.ConvertStringToUint(sid)
	if err != nil {
		response.FailedWithMessage(response.ParamErrorMessage)
		return
	}

	// 查询数据
	var areas []model.Area
	err = common.DB.Where("city_id = ?", uid).Find(&areas).Error
	if err != nil {
		response.FailedWithMessage("获取行政区信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": areas,
	})
}

// 根据区获取街道
func StreetListByAreaIdHandler(ctx *gin.Context) {
	// 参数处理
	sid := ctx.Param("areaId")
	uid, err := utils.ConvertStringToUint(sid)
	if err != nil {
		response.FailedWithMessage(response.ParamErrorMessage)
		return
	}

	// 查询数据
	var streets []model.Street
	err = common.DB.Where("area_id = ?", uid).Find(&streets).Error
	if err != nil {
		response.FailedWithMessage("获取街道信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": streets,
	})
}
