package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/ckeyer/puppy/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd())
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info("start.")
			cfg, err := config.LoadFile(configfile)
			if err != nil {
				logrus.Fatalln("load config failed, %s", err)
				return
			}

			logrus.Debugf("load config: %+v", cfg)
		},
	}

	return cmd
}
