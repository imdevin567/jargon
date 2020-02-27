package cmd

import (
	"github.com/imdevin567/jargon/bus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  "",
	Run:   run,
}

func init() {
	jargonCmd.AddCommand(runCmd)
}

func run(_ *cobra.Command, _ []string) {
	bus.Instance.Start()
}
