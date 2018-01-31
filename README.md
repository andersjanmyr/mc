# mc - CLI for Memcached

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

Flags:
      --config string   config file (default is $HOME/.mc.yaml)
  -h, --help            help for mc
      --host string     server hostname (default is localhost) (default "h")
      --port string     server port (default is 11211) (default "p")

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


