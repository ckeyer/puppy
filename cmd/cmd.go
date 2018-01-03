package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/spf13/cobra"
)

var (
	configfile string
	debug      bool

	rootCmd = &cobra.Command{
		Use:   "puppy",
		Short: "",
		Annotations: map[string]string{
			"a": "aae",
			"b": "bb",
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
			logrus.Debugf("start puppy.")
			logrus.Infof("debug?: %v", debug)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configfile, "config", "c", "~/.puppy.conf", "config file path.")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "debug")
}

func Execute() {
	rootCmd.Execute()
}
