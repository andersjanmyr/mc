package cmd

import (
	"fmt"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <key> [value]",
	Short: "Adds a key and value, if it doesn't already exist.",
	Long: `Adds a key and value, if it doesn't already exist.
Value can come from command line, -f <filename> or stdin.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		item, err := getItem(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		memcached := mc.Connect()
		err = memcached.Add(item)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.Flags().Int32VarP(&expiration, "expiration", "e", 0, "Expiration time for this key")
	addCmd.Flags().StringVarP(&file, "file", "f", "", "Filename containing the value")
	rootCmd.AddCommand(addCmd)
}
