package common

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 系统配置
var (
	Version     = "1.0"                 // 当前版本
	ConfigFile  = "config/default.yaml" // 默认配置文件
	VersionFile = "config/version"      // 版本文件
)

// 全局工具
var (
	DB    *gorm.DB      // 数据库连接
	Cache *redis.Client // 缓存连接
)
