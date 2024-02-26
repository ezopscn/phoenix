package cmd

import (
	"github.com/spf13/cobra"
	"phoenix/initialize"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(tableCmd)
	migrateCmd.AddCommand(dataCmd)
}

// 迁移命令
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate service database",
}

// 迁移表
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Migrate service table to MySQL",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateTable()
	},
}

// 迁移数据
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Migrate service data to MySQL",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateData()
	},
}
