package cmd

import (
	"fmt"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

var deleteAllCmd = &cobra.Command{
	Use:   "deleteall",
	Short: "Deletes all values from memcached",
	Long:  `Deletes all values from memcached`,
	Run: func(cmd *cobra.Command, args []string) {
		memcached := mc.Connect()
		err := memcached.DeleteAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteAllCmd)
}
