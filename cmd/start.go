package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stetsd/micro-brick/internal/app"
)

var port string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start http server",
	Long:  `start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		server := app.NewServer()
		server.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&port, "port", "p", "", "server port")
}
