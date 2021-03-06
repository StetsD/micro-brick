package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stetsd/micro-brick/config"
	"github.com/stetsd/micro-brick/internal/app"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start http server",
	Long:  `start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.EnvParseToConfigMap()

		if err != nil {
			panic(err)
		}

		server := app.NewServer(config)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
