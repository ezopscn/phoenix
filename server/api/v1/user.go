package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/dto"
	"phoenix/model"
	"phoenix/pkg/logx"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
	"strings"
)

// 获取用户总量处理函数
func GetUserCountHandler(ctx *gin.Context) {
	var count int64
	err := common.DB.Model(&model.User{}).Count(&count).Error
	if err != nil {
		response.FailedWithMessage("查询用户总数失败")
		return
	}
	response.SuccessWithData(gin.H{
		"count": count,
	})
}

// 获取用户列表处理函数
func GetUserListHandler(ctx *gin.Context) {
	// 获取请求参数
	var req dto.UserListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 用户列表
	var users []model.User

	// 查询模板
	dbt := common.DB.
		Preload("Role").
		Preload("Department").
		Preload("OfficeProvince").
		Preload("OfficeCity").
		Preload("OfficeArea").
		Preload("OfficeStreet").
		Preload("NativeProvince").
		Preload("NativeCity")

	// 查询条件
	// ENName
	if en_name := strings.TrimSpace(req.ENName); en_name != "" {
		dbt = dbt.Where("en_name LIKE ?", "%"+en_name+"%")
	}

	// CNName
	if cn_name := strings.TrimSpace(req.CNName); cn_name != "" {
		dbt = dbt.Where("cn_name LIKE ?", "%"+cn_name+"%")
	}

	// JobId
	if job_id := strings.TrimSpace(req.JobId); job_id != "" {
		dbt = dbt.Where("job_id LIKE ?", "%"+job_id+"%")
	}

	// JobName
	if job_name := strings.TrimSpace(req.JobName); job_name != "" {
		dbt = dbt.Where("job_name LIKE ?", "%"+job_name+"%")
	}

	// Phone
	if phone := strings.TrimSpace(req.Phone); phone != "" {
		dbt = dbt.Where("phone LIKE ?", "%"+phone+"%")
	}

	// Email
	if email := strings.TrimSpace(req.Email); email != "" {
		dbt = dbt.Where("email LIKE ?", "%"+email+"%")
	}

	// Gender
	if req.Gender != nil {
		dbt = dbt.Where("gender = ?", req.Gender)
	}

	// DepartmentId
	if req.DepartmentId != 0 {
		dbt = dbt.Where("department_id = ?", req.DepartmentId)
	}

	// OfficeProvinceId
	if req.OfficeProvinceId != 0 {
		dbt = dbt.Where("office_province_id = ?", req.OfficeProvinceId)
	}

	// OfficeCityId
	if req.OfficeCityId != 0 {
		dbt = dbt.Where("office_city_id = ?", req.OfficeCityId)
	}

	// OfficeAreaId
	if req.OfficeAreaId != 0 {
		dbt = dbt.Where("office_area_id = ?", req.OfficeAreaId)
	}

	// OfficeStreetId
	if req.OfficeStreetId != 0 {
		dbt = dbt.Where("office_street_id = ?", req.OfficeStreetId)
	}

	// OfficeAddress
	if office_address := strings.TrimSpace(req.OfficeAddress); office_address != "" {
		dbt = dbt.Where("office_address LIKE ?", "%"+office_address+"%")
	}

	// NativeProvinceId
	if req.NativeProvinceId != 0 {
		dbt = dbt.Where("native_province_id = ?", req.NativeProvinceId)
	}

	// NativeCityId
	if req.NativeCityId != 0 {
		dbt = dbt.Where("native_city_id = ?", req.NativeCityId)
	}

	// CreatorId
	if req.CreatorId != 0 {
		dbt = dbt.Where("creator_id = ?", req.CreatorId)
	}

	// Status
	if req.Status != nil {
		dbt = dbt.Where("status = ?", req.Status)
	}

	// RoleId
	if req.RoleId != 0 {
		dbt = dbt.Where("role_id = ?", req.RoleId)
	}

	// 判断是否分页
	var err error
	if !req.NoPagination {
		// 统计记录数量
		dbt.Find(&model.User{}).Count(&req.Total)
		// 获取偏移和限制
		limit, offset := req.GetLimitAndOffset()
		err = dbt.Limit(limit).Offset(offset).Find(&users).Error
	} else {
		err = dbt.Find(&users).Error
	}

	// 判断查询是否出错
	if err != nil {
		logx.ERROR("查询用户列表失败,", err.Error())
		response.FailedWithMessage("查询用户列表失败")
		return
	}

	// 获取当前用户的工号和角色关键字
	currentUserJobId, currentUserRoleKeyword, err := utils.GetJobIdAndRoleKeywordFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取当前用户的信息失败")
		return
	}

	// 当用户开启隐藏手机号，且当前用户不是管理员，信息获取的也不是自己，则需要对手机号加密
	if currentUserRoleKeyword != common.SuperAdminRoleKeyword {
		for idx, user := range users {
			if *user.ShowPhone == common.False {
				if currentUserJobId != user.JobId {
					users[idx].Phone = utils.MaskPhone(user.Phone)
				}
			}
		}
	}

	// 响应数据
	response.SuccessWithData(dto.PageData{
		Page: req.Page,
		List: users,
	})
}

// 查询用户信息
func GetUserInfoByJobId(ctx *gin.Context, jobId string) (user model.User, err error) {
	// 获取当前用户的工号和角色关键字
	currentUserJobId, currentUserRoleKeyword, err := utils.GetJobIdAndRoleKeywordFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取当前用户的信息失败")
		return
	}

	// 如果 jobId 为空，则表示当前用户
	if jobId == "" {
		jobId = currentUserJobId
	}

	// 查询用户信息
	err = common.DB.
		Preload("Role").
		Preload("Department").
		Preload("OfficeProvince").
		Preload("OfficeCity").
		Preload("OfficeArea").
		Preload("OfficeStreet").
		Preload("NativeProvince").
		Preload("NativeCity").
		Where("job_id = ?", jobId).First(&user).Error
	if err != nil {
		logx.ERROR("查询用户信息失败:", err.Error())
		return user, fmt.Errorf("查询用户信息失败")
	}

	// 当用户开启隐藏手机号，且当前用户不是管理员，信息获取的也不是自己，则需要对手机号加密
	if *user.ShowPhone == common.True {
		if currentUserRoleKeyword != common.SuperAdminRoleKeyword {
			if currentUserJobId != jobId {
				user.Phone = utils.MaskPhone(user.Phone)
			}
		}
	}

	return
}

// 获取当前用户信息处理函数
func GetCurrentUserInfoHandler(ctx *gin.Context) {
	// 获取用户信息
	user, err := GetUserInfoByJobId(ctx, "")
	if err != nil {
		response.FailedWithMessage("查询用户的信息失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"info": user,
	})
}

// 获取指定用户的用户信息处理函数
func GetSpecifyUserInfoHandler(ctx *gin.Context) {
	// 获取 URI 参数，并验证合法性
	jobId := ctx.Param("jobId")
	if !utils.IsJobId(jobId) {
		response.FailedWithMessage("查询用户的工号不合法")
		return
	}

	// 获取用户信息
	user, err := GetUserInfoByJobId(ctx, jobId)
	if err != nil {
		response.FailedWithMessage("查询用户的信息失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"info": user,
	})
}
