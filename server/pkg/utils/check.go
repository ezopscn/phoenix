package utils

import (
	"net"
	"regexp"
)

// 判断 IP 地址是否合法
func IsIPAddress(ip string) bool {
	result := net.ParseIP(ip)
	return result != nil
}

// 判断端口是否合法
func IsPort(port string) bool {
	_, err := net.ResolveTCPAddr("tcp", ":"+port)
	return err == nil
}

// 验证邮箱合法性
func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 验证手机号合法性
func IsPhone(phone string) bool {
	pattern := `^(13[0-9]|14[5-9]|15[0-3,5-9]|16[6]|17[0-8]|18[0-9]|19[8,9])\d{8}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

// 判断 JobId 合法性
// 只能包含大小写字母、数字和中横线
// 不能以数字或中横线开头
// 不能以中横线结尾
// 长度在 4-15 位
func IsJobId(jobId string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9-]{2,13}[a-zA-Z0-9]$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(jobId)
}

// 判断 RoleKeyword 合法性：
// 只能包含小写字母、数字和中横线
// 不能以数字或中横线开头
// 不能以中横线结尾
// 长度在 2-20 位
func IsRoleKeyword(roleKeyword string) bool {
	pattern := `^[a-z][a-z0-9-]{1,18}[a-z0-9]$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(roleKeyword)
}

// 隐藏手机号中间4位
func MaskPhone(phone string) string {
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}
