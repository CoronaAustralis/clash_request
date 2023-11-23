package cmd

import (
	"clash_request/server"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server at the specified address",
	Long: `Start the server at the specified address
the address is defined in config file`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Server()
	},
}

func init(){
	rootCmd.AddCommand(startCmd)
}
