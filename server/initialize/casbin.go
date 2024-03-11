package initialize

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"phoenix/common"
	"phoenix/pkg/logx"
)

// 初始化 Casbin 配置
func Casbin() {
	// 初始化数据库适配器
	adapter, err := gormadapter.NewAdapterByDBUseTableName(common.DB, "", "casbin_rule")
	if err != nil {
		logx.ERROR("Casbin连接初始失败")
		panic(err)
	}

	// 读取 embed 打包的配置
	bs, err := common.FS.ReadFile(common.CasbinFile)
	if err != nil {
		logx.ERROR("Casbin配置读取失败")
		panic(err)
	}

	// 从字符串中加载配置
	c := string(bs[:])
	m, err := casbinmodel.NewModelFromString(c)
	if err != nil {
		logx.ERROR("Casbin配置加载失败")
		panic(err)
	}

	// 读取配置文件
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		logx.ERROR("Casbin创建实例失败")
		panic(err)
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		logx.ERROR("Casbin策略加载失败")
		panic(err)
	}

	// 配置全局
	common.CasbinEnforcer = enforcer
	logx.SYSTEM("Casbin策略初始完成")
}
