package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
)

// 获取所有省份
func ProvinceListHandler(ctx *gin.Context) {
	var provinces []model.Province
	err := common.DB.Preload("Cities").Preload("Cities.Areas").Find(&provinces).Error
	if err != nil {
		response.FailedWithMessage("获取省份信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": provinces,
	})
}
