package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix/common"
	"phoenix/pkg/response"
	"phoenix/pkg/utils"
	"strings"
)

// Casbin 中间件
func Casbin(ctx *gin.Context) {
	// 获取用户角色关键字
	roleKeyword, err := utils.GetStringFromContext(ctx, "RoleKeyword")
	if err != nil || !utils.IsRoleKeyword(roleKeyword) {
		response.FailedWithMessage("获取用户角色信息失败")
		ctx.Abort()
		return
	}

	// sub，角色关键字
	sub := roleKeyword
	// obj，除去前缀和参数后的 URI，避免因为前缀更新了需要修改数据
	obj := strings.TrimPrefix(ctx.Request.RequestURI, common.Config.System.ApiPrefix+"/"+common.Config.System.ApiVersion)
	obj = strings.Split(obj, "?")[0]
	fmt.Println(obj)
	// act，请求方式
	act := ctx.Request.Method

	// 校验数据
	pass, _ := common.CasbinEnforcer.Enforce(sub, obj, act)
	if !pass {
		response.FailedWithMessage("校验用户角色权限失败")
		ctx.Abort()
		return
	}
	ctx.Next()
}
