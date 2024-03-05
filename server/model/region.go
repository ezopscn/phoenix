package model

// 数据来源：https://github.com/modood/Administrative-divisions-of-China
// 省
type Province struct {
	Id      uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name    string   `gorm:"comment:省份名称" json:"name"`
	Cities  []City   `json:"cities,omitempty"`
	Areas   []Area   `json:"areas,omitempty"`
	Streets []Street `json:"streets,omitempty"`
}

func (Province) TableName() string {
	return "province"
}

// 市
type City struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:市名称" json:"name"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	Areas      []Area   `json:"areas,omitempty"`
	Streets    []Street `json:"streets,omitempty"`
}

func (City) TableName() string {
	return "city"
}

// 区
type Area struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:区名称" json:"name"`
	CityId     uint     `gorm:"comment:市id" json:"city_id"`
	City       City     `gorm:"foreignKey:CityId;comment:市" json:"-"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	Streets    []Street `json:"streets,omitempty"`
}

func (Area) TableName() string {
	return "area"
}

// 街道
type Street struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:街道名称" json:"name"`
	AreaId     uint     `gorm:"comment:区id" json:"area_id"`
	Area       Area     `gorm:"foreignKey:CityId;comment:区" json:"-"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	CityId     uint     `gorm:"comment:市id" json:"city_id"`
	City       City     `gorm:"foreignKey:CityId;comment:市" json:"-"`
}

func (Street) TableName() string {
	return "street"
}
