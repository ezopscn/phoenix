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

	// 基础路由组
	brg := r.Group(baseApiPrefix)
	{
		routes.PublicRoutes(brg, auth) // 开放路由组
		routes.LoginRoutes(brg, auth)  // 登录路由组
		routes.RegionRoutes(brg, auth) // 地区路由组
		routes.UserRoutes(brg, auth)   // 用户路由组
		routes.RoleRoutes(brg, auth)   // 角色路由组
		routes.MenuRoutes(brg, auth)   // 菜单路由组
	}

	logx.SYSTEM("路由列表初始化完成")
	return r
}
