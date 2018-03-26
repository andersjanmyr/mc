package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <key(s)>",
	Short: "Gets one of more values",
	Long: `Gets one of more values.
	Keys should be comma separated without spaces. If no keys have values, get will exit with an error status. Multiple values will be returned one on each line in the order the keys were given. Missing keys will have the value key:none`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		memcached := mc.Connect()
		keys := strings.Split(args[0], ",")
		if len(keys) == 1 {
			it, err := memcached.Get(keys[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", string(it.Value))
		} else {
			results, err := memcached.GetMulti(keys)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if len(results) == 0 {
				fmt.Fprintln(os.Stderr, "No keys have values")
				os.Exit(1)
			}
			for _, k := range keys {
				it, ok := results[k]
				if !ok {
					fmt.Printf("%s:none\n", k)
				} else {
					fmt.Printf("%s\n", string(it.Value))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
