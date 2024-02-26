package initialize

import (
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/middleware"
	"phoenix/pkg/logx"
	"phoenix/routes"
)

// 路由初始化
func Router() *gin.Engine {
	// 创建一个没中间件的路由引擎
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog) // 访问日志中间件
	r.Use(middleware.Cors)      // 跨域访问中间件
	r.Use(middleware.Exception) // 异常捕获中间件

	// 路由组
	prg := r.Group(common.Config.System.ApiPrefix + "/" + common.Config.System.ApiVersion)
	routes.PublicRoutes(prg) // 开放路由组

	logx.SYSTEM("路由列表初始化完成")
	return r
}
