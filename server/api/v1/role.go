package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/dto"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
	"strings"
)

// 获取所有角色列表处理函数
func GetRoleListHandler(ctx *gin.Context) {
	// 获取请求参数
	var req dto.RoleListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 查询模板
	dbt := common.DB

	// 判断参数
	// Name
	if name := strings.TrimSpace(req.Name); name != "" {
		dbt = dbt.Where("name LIKE ?", "%"+name+"%")
	}

	// Keyword
	if keyword := strings.TrimSpace(req.Keyword); keyword != "" {
		dbt = dbt.Where("keyword LIKE ?", "%"+keyword+"%")
	}

	// Description
	if description := strings.TrimSpace(req.Description); description != "" {
		dbt = dbt.Where("description LIKE ?", "%"+description+"%")
	}

	// 查询数据
	var roles []model.Role
	err := dbt.Find(&roles).Error
	if err != nil {
		response.FailedWithMessage("查询角色列表失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"list": roles,
	})
}

// 获取指定角色详细信息处理函数
func GetRoleInfoByRoleKeywordHandler(ctx *gin.Context) {
	// 获取 URI 参数，并验证合法性
	roleKeyword := ctx.Param("roleKeyword")
	if !utils.IsRoleKeyword(roleKeyword) {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 查询数据
	var role model.Role
	err := common.DB.Where("keyword = ?", roleKeyword).First(&role).Error
	if err != nil {
		response.FailedWithMessage("查询角色信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"info": role,
	})
}
