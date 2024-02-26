package data

import (
	"errors"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"phoenix/common"
	"phoenix/model"
	"phoenix/pkg/utils"
)

// 初始化密码
var password = "ezops.cn"

// 用户初始化数据
var users = []model.User{
	{
		BaseModel:        model.BaseModel{Id: 1},
		Username:         "admin",
		ENName:           "administrator",
		CNName:           "超管",
		Mobile:           "18888888888",
		Email:            "admin@ezops.cn",
		Password:         utils.CryptoPassword(password),
		JobId:            "ez000001",
		JobName:          "高级运维工程师",
		JoinTime:         carbon.Now(),
		DepartmentId:     109400,
		OfficeProvinceId: 44,
		OfficeCityId:     4403,
		OfficeAreaId:     440304,
		OfficeStreetId:   440304005,
		OfficeAddress:    "下沙社区",
		OfficeStation:    "T1-9F-A-11",
		NativeProvinceId: 51,
		NativeCityId:     5105,
		Gender:           1,
		Avatar:           "/images/avatar/default.png",
		Birthday:         carbon.Now(),
		CreatorId:        0,
		RoleId:           1,
	},
	{
		BaseModel:        model.BaseModel{Id: 2},
		Username:         "guest",
		ENName:           "guest",
		CNName:           "访客",
		Mobile:           "19999999999",
		Email:            "guest@ezops.cn",
		Password:         utils.CryptoPassword(password),
		JobId:            "ez000002",
		JobName:          "访客",
		JoinTime:         carbon.Now(),
		DepartmentId:     109400,
		OfficeProvinceId: 44,
		OfficeCityId:     4403,
		OfficeAreaId:     440304,
		OfficeStreetId:   440304005,
		OfficeAddress:    "下沙社区",
		OfficeStation:    "T1-9F-A-12",
		NativeProvinceId: 51,
		NativeCityId:     5105,
		Gender:           2,
		Avatar:           "/images/avatar/default1.png",
		Birthday:         carbon.Now(),
		CreatorId:        0,
		RoleId:           2,
	},
}

// 用户数据初始化
func UserInit() {
	for _, item := range users {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.User
		err := common.DB.Where("id = ? OR username = ? OR mobile = ? OR email = ? OR job_id = ?",
			item.Id,
			item.Username,
			item.Mobile,
			item.Email,
			item.JobId).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
