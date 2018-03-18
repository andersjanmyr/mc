package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a key and value if it doesn't exist",
	Long:  `Adds a key and value if it doesn't exist`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		filename, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		var value []byte
		if len(args) == 2 {
			value = []byte(args[1])
		} else if filename != "" {
			value, err = ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Value or filename (-f) are required")
			os.Exit(1)
		}

		exp, err := cmd.Flags().GetInt32("expiration")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		memcached := mc.Connect()
		err = memcached.Add(&memcache.Item{
			Key:        key,
			Value:      value,
			Expiration: exp,
		})
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
