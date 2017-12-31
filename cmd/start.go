package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(startCmd())
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := viper.New()
			cfg.AddConfigPath(configfile)
			logrus.Infof("%+v", cfg.AllKeys())
		},
	}

	return cmd
}
