package cmd

import (
	"fmt"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

var touchCmd = &cobra.Command{
	Use:   "touch <key>",
	Short: "Touches a key (updates its expiration time)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		exp, err := cmd.Flags().GetInt32("expiration")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		memcached := mc.Connect()
		err = memcached.Touch(key, exp)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	touchCmd.Flags().Int32VarP(&expiration, "expiration", "e", 0, "Expiration time for this key")
	rootCmd.AddCommand(touchCmd)
}
