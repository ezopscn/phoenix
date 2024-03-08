package model

// 接口模型
type API struct {
	BaseModel
	API      string `gorm:"uniqueIndex:uidx_api;comment:接口URI" json:"api"`
	Method   string `gorm:"comment:请求方法" json:"method"`
	Name     string `gorm:"uniqueIndex:uidx_name;comment:接口名称" json:"name"`
	ParentId uint   `gorm:"comment:父接口Id" json:"parent_id"`
	Children []API  `gorm:"-" json:"children"`
}

// 表名称
func (API) TableName() string {
	return "api"
}
