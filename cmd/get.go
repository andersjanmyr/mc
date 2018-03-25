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
	Use:   "get <key>",
	Short: "Gets one of more values",
	Long: `Gets one of more values.
	Keys should be comma separated without spaces.
	Values will be returned one on each line in the order the keys were given.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		memcached := mc.Connect()
		keys := strings.Split(args[0], ",")
		results, err := memcached.GetMulti(keys)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, k := range keys {
			it, ok := results[k]
			if !ok {
				fmt.Printf("None\n")
			} else {
				fmt.Printf("%s\n", string(it.Value))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
