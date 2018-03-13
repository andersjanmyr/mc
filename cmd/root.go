package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	host    string
	port    string
	verbose bool
)

func log(args ...interface{}) {
	if !verbose {
		return
	}
	fmt.Fprintln(os.Stderr, args...)
}

var rootCmd = &cobra.Command{
	Use:   "mc",
	Short: "A memcached CLI client",
	Long: `mc is a command line client for memcached it supports the usual
commands, such as get, set, etc.`,
	ValidArgs: []string{"get", "set", "delete", "deleteall", "help",
		"version", "completion"},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mc.yaml)")
	rootCmd.PersistentFlags().StringVarP(&host, "server", "s", "localhost", "server hostname")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "11211", "server port")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("server", "localhost")
	viper.SetDefault("port", "11211")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".mc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log("Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		fmt.Println(err)
		os.Exit(1)
	}
}
