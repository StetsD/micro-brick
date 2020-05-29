package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stetsd/micro-brick/internal/app"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start http server",
	Long:  `start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
