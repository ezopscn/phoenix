package common

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 配置
var (
	Version                    = "1.0"                 // 当前版本
	ConfigFile                 = "config/default.yaml" // 默认配置文件
	VersionFile                = "config/version"      // 版本文件
	SuperAdminRoleKeyword      = "administrator"       // 系统超级管理员角色关键字，系统预留
	DefaultPageSize       uint = 1                     // 默认每页显示的数据量
	MaxPageSize           uint = 100                   // 每次请求最大的数据量
)

// 全局工具
var (
	DB    *gorm.DB      // 数据库连接
	Cache *redis.Client // 缓存连接
)
