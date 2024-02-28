package v1

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
)

// 用户列表
func UserListHandler(ctx *gin.Context) {
	var users []model.User
	err := common.DB.
		Preload("Role").
		Preload("Department").
		Preload("OfficeProvince").
		Preload("OfficeCity").
		Preload("OfficeArea").
		Preload("OfficeStreet").
		Preload("NativeProvince").
		Preload("NativeCity").
		Find(&users).Error
	if err != nil {
		response.FailedWithMessage("查询用户列表失败")
		return
	}

	// 处理手机号
	for idx, user := range users {
		if *user.ShowPhone == common.False {
			users[idx].Phone = utils.MaskPhone(user.Phone)
		}
	}

	response.SuccessWithData(map[string]interface{}{
		"list": users,
	})
}
