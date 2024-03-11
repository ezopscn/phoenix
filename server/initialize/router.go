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
	r.Use(middleware.AccessLog)       // 访问日志中间件
	r.Use(middleware.Cors)            // 跨域访问中间件
	r.Use(middleware.Exception)       // 异常捕获中间件
	auth, err := middleware.JWTAuth() // JWT认证中间件
	if err != nil {
		panic(err)
	}

	// 通用路由前缀
	baseApiPrefix := common.Config.System.ApiPrefix + "/" + common.Config.System.ApiVersion

	// 开放路由组
	prg := r.Group(baseApiPrefix)
	routes.PublicRoutes(prg, auth)

	// 认证路由组
	arg := r.Group(baseApiPrefix)
	arg.Use(auth.MiddlewareFunc())
	routes.AuthRoutes(arg, auth)

	// 鉴权路由组
	crg := r.Group(baseApiPrefix)
	crg.Use(auth.MiddlewareFunc())
	crg.Use(middleware.Casbin)
	routes.CasbinRouter(crg, auth)

	logx.SYSTEM("路由列表初始化完成")
	return r
}
