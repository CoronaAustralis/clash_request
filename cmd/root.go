package cmd

import (
	"clash_request/cmd/flags"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clash_request",
	Short: "a clash yaml file subsrciption",
	Long:  `a clash subscription to get clash yaml file, and will expose a port to let other request to get yaml file `,
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&flags.Config, "./config.json", "", "config file path")
	rootCmd.Flags().IntVar(&flags.Port, "port", 7778, "set port")
}

