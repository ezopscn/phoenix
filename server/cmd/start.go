package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"phoenix/common"
	"phoenix/initialize"
	"phoenix/pkg/logx"
	"phoenix/pkg/utils"
	"time"
)

func init() {
	rootCmd.AddCommand(startCmd)
	// 指定配置文件参数
	startCmd.Flags().StringVarP(&common.ConfigFile, "config", "f", common.ConfigFile, "specify run config for service")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service with some flags",
	Run: func(cmd *cobra.Command, args []string) {
		// Logo
		fmt.Println(common.Logo)

		// 配置文件初始化
		initialize.Config()

		// MySQL 连接初始化
		initialize.MySQL()

		// Redis 连接初始化
		initialize.Redis()

		// 初始化 Casbin 鉴权
		initialize.Casbin()

		// 路由初始化
		r := initialize.Router()

		// 判断参数是否合法
		if !utils.IsIPAddress(common.Config.System.Listen) {
			logx.ERROR("服务监听地址不合法")
			return
		}

		// 检测端口是否合法
		if !utils.IsPort(common.Config.System.Port) {
			logx.ERROR("服务监听端口不合法")
			return
		}

		// 监听地址
		listenAddress := fmt.Sprintf("%s:%s", common.Config.System.Listen, common.Config.System.Port)
		logx.SYSTEM("服务启动监听的地址:", listenAddress)

		// 配置服务
		server := http.Server{
			Addr:    listenAddress,
			Handler: r,
		}

		// 启动服务
		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				logx.ERROR(err)
				panic(err)
			}
		}()

		// 接收优雅关闭信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		// 等待5秒然后停止服务
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			logx.ERROR(err)
			panic(err)
		}
		logx.SYSTEM("服务正常的停止完成")
	},
}
