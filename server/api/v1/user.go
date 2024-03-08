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
		response.FailedWithMessage(response.ParamErrorMessage)
		return
	}

	// 获取当前用户的角色关键字
	roleKeyword, err := utils.GetStringFromContext(ctx, "RoleKeyword")
	if err != nil {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 如果不是管理员，需要处理手机号显示问题
	if roleKeyword != common.SuperAdminRoleKeyword {
		for idx, user := range users {
			if *user.ShowPhone == common.False {
				users[idx].Phone = utils.MaskPhone(user.Phone)
			}
		}
	}

	// 响应数据
	response.SuccessWithData(dto.PageData{
		Page: req.Page,
		List: users,
	})
}

// 获取当前用户信息处理函数
func GetCurrentUserInfoHandler(ctx *gin.Context) {
	// 获取当前用户的 JobId
	jobId, err := utils.GetStringFromContext(ctx, "JobId")
	if err != nil || !utils.IsJobId(jobId) {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 查询当前用户的用户信息
	var user model.User
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
		response.FailedWithMessage("查询当前用户的用户信息失败")
		return
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"info": user,
	})
}

// 获取指定用户的用户信息处理函数
// 支持参数：dto.UserInfoRequest
func GetSpecifyUserInfoHandler(ctx *gin.Context) {
	// 获取当前用户的 JobId
	jobId, err := utils.GetStringFromContext(ctx, "JobId")
	if err != nil || !utils.IsJobId(jobId) {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 获取当前用户的角色关键字
	roleKeyword, err := utils.GetStringFromContext(ctx, "RoleKeyword")
	if err != nil || !utils.IsRoleKeyword(roleKeyword) {
		response.FailedWithCodeAndMessage(response.Forbidden, response.ForbiddenMessage)
		return
	}

	// 获取 URI 参数，并验证合法性
	sid := ctx.Param("jobId")
	if !utils.IsJobId(sid) {
		response.FailedWithCodeAndMessage(response.ParamError, response.ParamErrorMessage)
		return
	}

	// 查询当前用户的用户信息
	var user model.User
	err = common.DB.
		Preload("Role").
		Preload("Department").
		Preload("OfficeProvince").
		Preload("OfficeCity").
		Preload("OfficeArea").
		Preload("OfficeStreet").
		Preload("NativeProvince").
		Preload("NativeCity").
		Where("job_id = ?", sid).First(&user).Error

	if err != nil {
		response.FailedWithMessage("查询指定用户的用户信息失败")
		return
	}

	// 不是管理员且查询的用户不是自己，则需要检查密码是否需要隐藏
	if roleKeyword != common.SuperAdminRoleKeyword && jobId != sid {
		if *user.ShowPhone == common.False {
			user.Phone = utils.MaskPhone(user.Phone)
		}
	}

	// 响应请求
	response.SuccessWithData(gin.H{
		"info": user,
	})
}
