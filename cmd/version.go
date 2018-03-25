package cmd

import (
	"fmt"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(mc.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
