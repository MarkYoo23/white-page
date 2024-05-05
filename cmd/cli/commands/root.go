package commands

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	RootCmd = &cobra.Command{
		Use:   "the-write-page-cli",
		Short: "The write page applications",
		Long:  ``,
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// do something
}
