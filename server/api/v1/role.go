package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
)

// 获取所有角色列表
func RoleListHandler(ctx *gin.Context) {
	var roles []model.Role
	err := common.DB.Find(&roles).Error
	if err != nil {
		response.FailedWithMessage("查询角色列表失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": roles,
	})
}

// 角色信息
func RoleInfoByIdHandler(ctx *gin.Context) {
	// URI 参数
	sid := ctx.Param("roleId")
	roleId, err := utils.ConvertStringToUint(sid)
	if err != nil || roleId == 0 {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 查询数据
	var role model.Role
	err = common.DB.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		response.FailedWithMessage("查询角色信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"info": role,
	})
}
