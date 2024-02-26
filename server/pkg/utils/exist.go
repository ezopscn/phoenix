package utils

import (
	"fmt"
	"os"
)

// 判断文件是否存在
func FileExists(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		if !stat.IsDir() {
			return true, nil
		}
		return false, fmt.Errorf("路径 %s 已经存在, 但它是一个目录", filename)
	}
	return false, fmt.Errorf("路径 %s 不存在", filename)
}

// 判断目录是否存储
func DirExists(dirname string) (bool, error) {
	stat, err := os.Stat(dirname)
	if !os.IsNotExist(err) {
		if stat.IsDir() {
			return true, nil
		}
		return false, fmt.Errorf("路径 %s 已经存在, 但它是一个文件", dirname)
	}
	return false, fmt.Errorf("路径 %s 不存在", dirname)
}
