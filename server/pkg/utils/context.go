package utils

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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
