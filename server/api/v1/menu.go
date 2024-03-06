package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
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

// 通过角色获取菜单列表
func GetMenuTreeByRoleId(roleId uint) (tree []model.Menu, err error) {
	var menus []model.Menu
	var role model.Role

	// 获取指定角色的菜单列表
	if utils.ContainsUint(common.AdminRoleIds, roleId) {
		// 管理员则获取所有的菜单
		err = common.DB.Find(&menus).Error
	} else {
		// 查询指定角色的菜单
		err = common.DB.Preload("Menus").Where("id = ?", roleId).First(&role).Error
		menus = role.Menus
	}

	if err != nil {
		return tree, fmt.Errorf("查询指定角色的菜单列表失败")
	}

	// 生成菜单树
	tree = GenerateMenuTree(0, menus)
	return
}

// 获取当前用户角色的菜单列表
func CurrentUserMenuListHandler(ctx *gin.Context) {
	// 当前用户的角色 Id
	roleId, err := utils.GetUintFromContext(ctx, "RoleId")
	if err != nil {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 生成菜单树
	menus, _ := GetMenuTreeByRoleId(roleId)
	response.SuccessWithData(gin.H{
		"list": menus,
	})
}

// 获取指定角色的菜单列表
func MenuListByRoleIdHandler(ctx *gin.Context) {
	// URI 参数
	sid := ctx.Param("roleId")
	roleId, err := utils.ConvertStringToUint(sid)
	if err != nil || roleId == 0 {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 生成菜单树
	menus, _ := GetMenuTreeByRoleId(roleId)
	response.SuccessWithData(gin.H{
		"list": menus,
	})
}
