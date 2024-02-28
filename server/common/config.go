package common

import "embed"

// 配置打包
var FS embed.FS

// 配置引用
var Config Configuration

// 配置结构体
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	MySQL  MySQLConfiguration  `mapstructure:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" json:"redis"`
	JWT    JWTConfiguration    `mapstructure:"jwt" json:"jwt"`
	Login  LoginConfiguration  `mapstructure:"login" json:"login"`
}

// 系统配置
type SystemConfiguration struct {
	Listen     string `mapstructure:"listen" json:"listen"`
	Port       string `mapstructure:"port" json:"port"`
	ApiPrefix  string `mapstructure:"api-prefix" json:"api_prefix"`
	ApiVersion string `mapstructure:"api-version" json:"api_version"`
	LogLevel   int    `mapstructure:"log-level" json:"log_level"`
}

// 数据库配置
type MySQLConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Charset      string `mapstructure:"charset" json:"charset"`
	Collation    string `mapstructure:"collation" json:"collation"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	ExtraParam   string `mapstructure:"extra-param" json:"extra_param"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"max_idle_time"`
}

// Redis 配置
type RedisConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     int    `mapstructure:"database" json:"database"`
	Password     string `mapstructure:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max_open_conns"`
	MinIdleConns int    `mapstructure:"min-idle-conns" json:"min_idle_conns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"max_idle_time"`
}

// JWT 配置
type JWTConfiguration struct {
	Realm   string `mapstructure:"realm" json:"realm"`
	Key     string `mapstructure:"key" json:"key"`
	Timeout int    `mapstructure:"timeout" json:"timeout"`
}

// 登录配置
type LoginConfiguration struct {
	WrongTimes     int  `mapstructure:"wrong-times" json:"wrong_times"`
	LockTime       int  `mapstructure:"lock-time" json:"lock_time"`
	MultiDevices   bool `mapstructure:"multi-devices" json:"multi_devices"`
	ResetTokenTime int  `mapstructure:"reset-token-time" json:"reset_token_time"`
}
