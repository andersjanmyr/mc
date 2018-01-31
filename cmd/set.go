package cmd

import (
	"fmt"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/rainycape/memcache"
	"github.com/spf13/cobra"
)

var expiration int32

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a key and value in memcached",
	Long:  `Sets a key and value in memcached`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		memcached, err := mc.Connect()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		exp, err := cmd.Flags().GetInt32("expiration")
		fmt.Println(exp)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = memcached.Set(&memcache.Item{
			Key:        args[0],
			Value:      []byte(args[1]),
			Expiration: exp,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	setCmd.Flags().Int32VarP(&expiration, "expiration", "e", 0, "Expiration time for this key")
	rootCmd.AddCommand(setCmd)
}
