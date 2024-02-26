package logx

import (
	"fmt"
	"phoenix/common"
	"runtime"
	"time"
)

// 日志级别
const (
	LOG_ACCESS  = iota // 访问级别
	LOG_DEBUG          // 调试级别（默认）
	LOG_INFO           // 消息级别
	LOG_WARNING        // 警告级别
	LOG_ERROR          // 错误级别
	LOG_SYSTEM         // 系统级别
)

// 日志级别对应的内容
const (
	LOG_ACCESS_CONTEN   = "ACCESS"
	LOG_DEBUG_CONTENT   = "DEBUG"   // 调试
	LOG_INFO_CONTENT    = "INFO"    // 消息
	LOG_WARNING_CONTENT = "WARNING" // 警告
	LOG_ERROR_CONTENT   = "ERROR"   // 错误
	LOG_SYSTEM_CONTENT  = "SYSTEM"  // 系统
)

// 日志级别和内容绑定
var LogLevelContent = map[int]string{
	LOG_ACCESS:  LOG_ACCESS_CONTEN,
	LOG_DEBUG:   LOG_DEBUG_CONTENT,
	LOG_INFO:    LOG_INFO_CONTENT,
	LOG_WARNING: LOG_WARNING_CONTENT,
	LOG_ERROR:   LOG_ERROR_CONTENT,
	LOG_SYSTEM:  LOG_SYSTEM_CONTENT,
}

// 日志打印
func PrintLog(level int, v ...interface{}) {
	if level >= common.Config.System.LogLevel {
		now := time.Now().Format(common.MsecTimeFormat)
		pc, _, line, _ := runtime.Caller(2)
		fn := runtime.FuncForPC(pc).Name()
		fmt.Print(fmt.Sprintf("%s\t%s\t%s:%d\t", now, LogLevelContent[level], fn, line), fmt.Sprintln(v...))
	}
}

// DEBUG 级别日志
func DEBUG(v ...interface{}) {
	PrintLog(LOG_DEBUG, v...)
}

// DEBUG 级别日志
func INFO(v ...interface{}) {
	PrintLog(LOG_INFO, v...)
}

// DEBUG 级别日志
func WARNING(v ...interface{}) {
	PrintLog(LOG_WARNING, v...)
}

// DEBUG 级别日志
func ERROR(v ...interface{}) {
	PrintLog(LOG_ERROR, v...)
}

// SYSTEM 级别日志
func SYSTEM(v ...interface{}) {
	PrintLog(LOG_SYSTEM, v...)
}

// ACCESS 请求访问日志单独的格式
func ACCESS(v ...interface{}) {
	now := time.Now().Format(common.MsecTimeFormat)
	fmt.Print(fmt.Sprintf("%s\t%s\t", now, LogLevelContent[LOG_ACCESS]), fmt.Sprintln(v...))
}
