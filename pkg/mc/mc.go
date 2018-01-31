package mc

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rainycape/memcache"
	"github.com/spf13/viper"
)

// Connect to the configured server
func Connect() (*memcache.Client, error) {
	url := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	memcached, err := memcache.New(url)
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	return memcached, nil
}
