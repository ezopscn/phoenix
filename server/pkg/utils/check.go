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
func IsMobile(mobile string) bool {
	pattern := `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}
