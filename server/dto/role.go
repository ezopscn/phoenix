package dto

// 角色列表请求（包含查询筛选条件）
type RoleListRequest struct {
	Name        string `json:"name" form:"name"`
	Keyword     string `json:"keyword" form:"keyword"`
	Description string `json:"description" form:"description"`
}
