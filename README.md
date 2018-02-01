# mc - CLI for Memcached

## Installation

`mc` is a single binary. Install it by right-clicking and `Save as...` or with
`curl`.

### Links

* [OS X](https://github.com/andersjanmyr/mc/releases/download/v1.0.0/mc-osx)
* [Linux](https://github.com/andersjanmyr/mc/releases/download/v1.0.0/mc-linux)
* [Windows](https://github.com/andersjanmyr/mc/releases/download/v1.0.0/mc.exe)

### Curl

```
# OS X
$ curl -L https://github.com/andersjanmyr/mc/releases/download/v1.0.0/mc-osx \
  > /usr/local/bin/mc

# Linux
$ curl -L https://github.com/andersjanmyr/mc/releases/download/v1.0.0/mc-linux \
  > /usr/local/bin/mc

# Make executable
$ chmod a+x /usr/local/bin/mc
```

## Usage

```
$ ./mc
mc is a command line client for memcached it supports the usual
commands, such as get, set, etc.

Usage:
  mc [command]

Available Commands:
  delete      Delete a value from memcached
  deleteall   Deletes all values from memcached
  get         Gets a value from memcached
  help        Help about any command
  set         Sets a key and value in memcached
  version     Print the version of mc

Flags:
      --config string   config file (default is $HOME/.mc.yaml)
  -h, --help            help for mc
  -p, --port string     server port (default "11211")
  -s, --server string   server hostname (default "localhost")

Use "mc [command] --help" for more information about a command.
```

### Get

```
$ ./mc get -h
Get a value from memcached by key

Usage:
  mc get [flags]

Flags:
  -h, --help   help for get
```

### Set

```
$ ./mc set -h
Sets a key and value in memcached

Usage:
  mc set [flags]

Flags:
  -e, --expiration int32   Expiration time for this key
  -f, --file string        Filename containing the value
  -h, --help               help for set
```

### Delete

```
$ ./mc delete -h
Delete a value from memcached by key

Usage:
  mc delete [flags]
```

### Delete All

```
$ ./mc deleteall -h
Deletes all values from memcached

Usage:
  mc deleteall [flags]
```


