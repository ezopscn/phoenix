package utils

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"phoenix/pkg/logx"
)

// 获取 Uint 属性
func GetUintFromContext(ctx *gin.Context, keyword string) (value uint, err error) {
	claims := jwt.ExtractClaims(ctx)
	v, _ := claims[keyword].(float64) // 注意客户端请求过来的 JSON 会变成 float64 类型
	if v == 0 {
		return value, fmt.Errorf("获取请求用户的%s失败", keyword)
	}
	value = uint(v)
	return
}

// 获取 String 属性
func GetStringFromContext(ctx *gin.Context, keyword string) (value string, err error) {
	claims := jwt.ExtractClaims(ctx)
	value, _ = claims[keyword].(string)
	return
}

// 获取 JobId
func GetJobIdFromContext(ctx *gin.Context) (jobId string, err error) {
	jobId, err = GetStringFromContext(ctx, "JobId")
	if !IsJobId(jobId) {
		err = errors.New("获取到的员工工号不合法")
	}
	if err != nil {
		logx.ERROR("获取当前用户的工号失败,", err.Error())
	}
	return
}

// 获取 RoleKeyword
func GetRoleKeywordFromContext(ctx *gin.Context) (roleKeyword string, err error) {
	roleKeyword, err = GetStringFromContext(ctx, "RoleKeyword")
	if !IsRoleKeyword(roleKeyword) {
		err = errors.New("获取到的角色关键字不合法")
	}
	if err != nil {
		logx.ERROR("获取当前用户的角色信息失败,", err.Error())
	}
	return
}

// JobId 和 RoleKeyword 都获取
func GetJobIdAndRoleKeywordFromContext(ctx *gin.Context) (jobId string, roleKeyword string, err error) {
	// JobId
	jobId, err = GetJobIdFromContext(ctx)
	if err != nil {
		return "", "", err
	}

	// RoleKeyword
	roleKeyword, err = GetRoleKeywordFromContext(ctx)
	if err != nil {
		return "", "", err
	}

	return
}
