package utils

import (
	"strconv"
)

// 判断字符串参数是不是数字，如果是，转换成 uint
func ConvertStringToUint(s string) (u uint, err error) {
	num, err := strconv.ParseUint(s, 10, 64)
	return uint(num), err
}
