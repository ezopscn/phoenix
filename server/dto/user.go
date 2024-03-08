package dto

// 用户列表请求（包含查询筛选条件）
type UserListRequest struct {
	ENName           string `json:"en_name" form:"en_name"`
	CNName           string `json:"cn_name" form:"cn_name"`
	JobId            string `json:"job_id" form:"job_id"`
	JobName          string `json:"job_name" form:"job_name"`
	Phone            string `json:"phone" form:"phone"`
	Email            string `json:"email" form:"email"`
	Gender           *uint  `json:"gender" form:"gender"`
	DepartmentId     uint   `json:"department_id" form:"department_id"`
	OfficeProvinceId uint   `json:"office_province_id" form:"office_province_id"`
	OfficeCityId     uint   `json:"office_city_id" form:"office_city_id"`
	OfficeAreaId     uint   `json:"office_area_id" form:"office_area_id"`
	OfficeStreetId   uint   `json:"office_street_id" form:"office_street_id"`
	OfficeAddress    string `json:"office_address" form:"office_address"`
	NativeProvinceId uint   `json:"native_province_id" form:"native_province_id"`
	NativeCityId     uint   `json:"native_city_id" form:"native_city_id"`
	CreatorId        uint   `json:"creator_id" form:"creator_id"`
	Status           *uint  `json:"status" form:"status"`
	RoleId           uint   `json:"role_id" form:"role_id"`
	Page
}

// 用户信息请求
type UserInfoRequest struct {
	JobId string `json:"job_id" form:"job_id"`
}
