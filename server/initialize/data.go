package initialize

import "phoenix/initialize/data"

// 数据初始化
func MigrateData() {
	data.RegionInit()     // 省市区街道数据
	data.DepartmentInit() // 部门数据
	data.RoleInit()       // 角色数据
	data.MenuInit()       // 菜单数据
	data.UserInit()       // 用户数据
}
