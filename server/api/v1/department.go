package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
)

// 寻找父级部门
func FindParentDepartments(dept model.Department, depts []model.Department) (parent model.Department, err error) {
	for _, d := range depts {
		if d.Id == dept.ParentId {
			d.Children = append(d.Children, dept)
			if d.ParentId != 0 {
				d, err = FindParentDepartments(d, depts)
			}
			parent = d
		}
	}
	return
}

// 根据部门 Id 获取部门树
func GetDepartmentTreeByDepartmentId(id uint) (dept model.Department, err error) {
	var departments []model.Department
	var department model.Department

	// 查询当前菜单
	err = common.DB.Where("id = ?", id).First(&department).Error
	if err != nil {
		return dept, fmt.Errorf("获取查询到部门信息失败")
	}

	// 查询所有菜单
	err = common.DB.Find(&departments).Error
	if err != nil {
		return dept, fmt.Errorf("获取所有的部门信息失败")
	}

	// 如果本身是最顶级部门，则不需要在查询
	if department.ParentId == 0 {
		return department, nil
	}

	// 查询父级部门
	return FindParentDepartments(department, departments)
}

// 获取当前用户部门信息
func GetCurrentUserDepartmentInfo(ctx *gin.Context) {
	// 获取当前用户的部门 Id
	id, err := utils.GetUintFromContext(ctx, "DepartmentId")
	if err != nil || id == 0 {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 获取部门信息
	dept, err := GetDepartmentTreeByDepartmentId(id)
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}
	response.SuccessWithData(gin.H{
		"info": dept,
	})
}
