package cmd

import (
	"fmt"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace <key> [value]",
	Short: "Replaces a key and value, if it already exists",
	Long: `Replaces a key and value, if it already exists.
Value can come from command line, -f <filename> or stdin.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		item, err := getItem(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		memcached := mc.Connect()
		err = memcached.Replace(item)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	replaceCmd.Flags().Int32VarP(&expiration, "expiration", "e", 0, "Expiration time for this key")
	replaceCmd.Flags().StringVarP(&file, "file", "f", "", "Filename containing the value")
	rootCmd.AddCommand(replaceCmd)
}
