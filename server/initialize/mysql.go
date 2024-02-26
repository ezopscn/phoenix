package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"phoenix/common"
	"phoenix/pkg/logx"
	"time"
)

// MySQL 连接初始化
func MySQL() {
	// 数据库连接串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&timeout=%dms&%s",
		common.Config.MySQL.Username,
		common.Config.MySQL.Password,
		common.Config.MySQL.Host,
		common.Config.MySQL.Port,
		common.Config.MySQL.Database,
		common.Config.MySQL.Charset,
		common.Config.MySQL.Collation,
		common.Config.MySQL.Timeout,
		common.Config.MySQL.ExtraParam)

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // 数据库连接字符串
		DefaultStringSize: 170, // varchar 默认长度，太长影响查询
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
			// TablePrefix:   "tb_", // 表名前缀
		},
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键
		IgnoreRelationshipsWhenMigrating:         false, // 开启会导致 many2many 的表创建失败
		QueryFields:                              true,  // 解决查询索引失效问题
	})

	// 错误处理
	if err != nil {
		logx.ERROR(err)
		panic("MySQL 连接初始失败")
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(common.Config.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(common.Config.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(common.Config.MySQL.MaxIdleTime) * time.Minute)

	// 设置全局数据库连接，方便后续使用
	common.DB = db
	logx.SYSTEM("MySQL 连接初始完成:", fmt.Sprintf("%s:%d/%s",
		common.Config.MySQL.Host,
		common.Config.MySQL.Port,
		common.Config.MySQL.Database))
}
