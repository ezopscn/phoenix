package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/dto"
	"phoenix/model"
	"phoenix/pkg/logx"
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

	// 判断是否分页
	var err error
	if !req.NoPagination {
		// 统计记录数量
		dbt.Find(&model.Role{}).Count(&req.Total)
		// 获取偏移和限制
		limit, offset := req.GetLimitAndOffset()
		err = dbt.Limit(limit).Offset(offset).Find(&roles).Error
	} else {
		err = dbt.Find(&roles).Error
	}

	if err != nil {
		response.FailedWithMessage("查询角色列表失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"list": roles,
	})
}

// 获取指定角色的详细信息
func GetRoleInfoByRoleKeyword(ctx *gin.Context, roleKeyword string) (role model.Role, err error) {
	// 当传递的 roleKeyword 为空，则表示查询当前用户的角色信息
	if roleKeyword == "" {
		roleKeyword, err = utils.GetRoleKeywordFromContext(ctx)
		if err != nil {
			return
		}
	}

	// 查询数据
	err = common.DB.Where("keyword = ?", roleKeyword).First(&role).Error
	if err != nil {
		logx.ERROR("查询角色信息失败,", err.Error())
	}
	return
}

// 获取当前用户角色详细信息处理函数
func GetCurrentUserRoleInfoHandler(ctx *gin.Context) {
	// 查询
	role, err := GetRoleInfoByRoleKeyword(ctx, "")
	if err != nil {
		response.FailedWithMessage("查询当前用户角色信息失败")
		return
	}

	// 响应
	response.SuccessWithData(gin.H{
		"info": role,
	})
}

// 获取指定角色详细信息处理函数
func GetSpecifyRoleInfoHandler(ctx *gin.Context) {
	// 获取 URI 参数，并验证合法性
	roleKeyword := ctx.Param("roleKeyword")
	if !utils.IsRoleKeyword(roleKeyword) {
		response.FailedWithMessage("查询的角色关键字不合法")
		return
	}

	// 查询
	role, err := GetRoleInfoByRoleKeyword(ctx, roleKeyword)
	if err != nil {
		response.FailedWithMessage("查询角色信息失败")
		return
	}

	// 响应
	response.SuccessWithData(gin.H{
		"info": role,
	})
}
