package data

import (
	"errors"
	"gorm.io/gorm"
	"phoenix/common"
	"phoenix/model"
)

// 角色数据
var roles = []model.Role{
	{
		BaseModel:   model.BaseModel{Id: 1},
		Name:        "超级管理员",
		Keyword:     "Administrator",
		Description: "超级管理员，最高权限",
	},
	{
		BaseModel:   model.BaseModel{Id: 2},
		Name:        "访客",
		Keyword:     "Guest",
		Description: "游客，最基础的浏览权限",
	},
}

// 角色数据初始化
func RoleInit() {
	for _, item := range roles {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.Role
		err := common.DB.Where("id = ? OR name = ?", item.Id, item.Name).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
