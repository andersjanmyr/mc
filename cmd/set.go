package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andersjanmyr/mc/pkg/mc"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var expiration int32
var file string

func getItem(cmd *cobra.Command, args []string) (*memcache.Item, error) {
	key := args[0]
	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get file argument")
	}
	var value []byte
	if len(args) == 2 {
		value = []byte(args[1])
	} else if filename != "" {
		value, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, errors.Wrapf(err, "Error reading file %s", filename)
		}
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			value, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				return nil, errors.Wrapf(err, "Error reading stdin %s", filename)
			}
		} else {
			return nil, errors.New("Value, stdin or filename (-f) are required")
		}
	}
	exp, err := cmd.Flags().GetInt32("expiration")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get expiration argument")
	}
	return &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: exp,
	}, nil
}

var setCmd = &cobra.Command{
	Use:   "set <key> [value]",
	Short: "Sets a key and value.",
	Long: `Sets a key and value.
Value can come from command line, -f <filename> or stdin.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		item, err := getItem(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		memcached := mc.Connect()
		err = memcached.Set(item)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	setCmd.Flags().Int32VarP(&expiration, "expiration", "e", 0, "Expiration time for this key")
	setCmd.Flags().StringVarP(&file, "file", "f", "", "Filename containing the value")
	rootCmd.AddCommand(setCmd)
}
