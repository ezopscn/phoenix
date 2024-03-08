package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/dto"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
	"strings"
)

// 菜单树生成
func GenerateMenuTree(parentId uint, menus []model.Menu) (tree []model.Menu) {
	for _, menu := range menus {
		if menu.ParentId == parentId {
			menu.Children = GenerateMenuTree(menu.Id, menus)
			tree = append(tree, menu)
		}
	}
	return
}

// 通过角色关键字获取菜单列表
func GetMenuTreeByRoleKeyword(roleKeyword string) (tree []model.Menu, err error) {
	var menus []model.Menu
	var role model.Role

	// 获取指定角色的菜单列表
	if roleKeyword == common.SuperAdminRoleKeyword {
		// 管理员则获取所有的菜单
		err = common.DB.Find(&menus).Error
	} else {
		// 查询指定角色的菜单
		err = common.DB.Preload("Menus").Where("keyword = ?", roleKeyword).First(&role).Error
		menus = role.Menus
	}

	if err != nil {
		return tree, fmt.Errorf("查询指定角色的菜单列表失败")
	}

	// 生成菜单树
	tree = GenerateMenuTree(0, menus)
	return
}

// 获取所有菜单列表处理函数
func GetMenuListHandler(ctx *gin.Context) {
	// 获取请求参数，并判断是否有参数
	var req dto.MenuListRequest
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

	// 查询数据
	var menus []model.Menu
	err := dbt.Find(&menus).Error
	if err != nil {
		response.FailedWithMessage("查询菜单列表失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"list": menus,
	})
}

// 获取指定角色的菜单列表处理函数
func GetSpecifyRoleKeywordMenuListHandler(ctx *gin.Context) {
	// 获取 URI 参数，并验证合法性
	roleKeyword := ctx.Param("roleKeyword")
	if !utils.IsRoleKeyword(roleKeyword) {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 生成菜单树
	menus, err := GetMenuTreeByRoleKeyword(roleKeyword)
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"list": menus,
	})
}

// 获取菜单的详细信息处理函数
func GetMenuInfoHandler(ctx *gin.Context) {
	// 获取 URI 参数，并验证合法性
	sid := ctx.Param("menuId")
	menuId, err := utils.ConvertStringToUint(sid)
	if err != nil || menuId == 0 {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 查询指定菜单
	var menu model.Menu
	err = common.DB.Where("id = ?", menuId).First(&menu).Error
	if err != nil {
		response.FailedWithMessage("查询菜单详细信息失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"info": menu,
	})
}
