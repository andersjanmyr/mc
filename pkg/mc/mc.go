package mc

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/viper"
)

// Connect to the configured server
func Connect() *memcache.Client {
	url := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	memcached := memcache.New(url)
	return memcached
}
