# ![mc-logo](./images/mc-small.png) mc - CLI for Memcached

## Installation

### Homebrew Installation on OS X

```
$ brew tap andersjanmyr/mc
$ brew install mc
```

### Other Installation

`mc` is a single binary. Install it by right-clicking and `Save as...` or with
`curl`.

### Links

* [OS X](https://github.com/andersjanmyr/mc/releases/download/v1.4.2/mc-osx)
* [Linux](https://github.com/andersjanmyr/mc/releases/download/v1.4.2/mc-linux)
* [Windows](https://github.com/andersjanmyr/mc/releases/download/v1.4.2/mc.exe)

### Curl

```
# OS X
$ curl -L https://github.com/andersjanmyr/mc/releases/download/v1.4.2/mc-osx \
  > /usr/local/bin/mc

# Linux
$ curl -L https://github.com/andersjanmyr/mc/releases/download/v1.4.2/mc-linux \
  > /usr/local/bin/mc

# Make executable
$ chmod a+x /usr/local/bin/mc
```

## Examples

```
mc get foo                 # Get value for key 'foo'
mc get foo,bar             # Get values for keys 'foo' and 'bar'
mc set foo bar             # Set value 'bar' for key 'foo'
mc set image -f image.png  # Set value for key 'image' to contents of file 'image.png'
echo -n 'bar' | mc set foo # Sets value 'bar' for key 'foo' from stdin
mc delete foo              # Delete key (and value) 'foo'
mc deleteall               # Delete all keys
mc add foo bar             # Add value 'bar' for key 'foo' if key doesn't exist
mc replace foo bar         # Replace value for key 'foo' with 'bar'
mc touch foo -e 30         # Set key 'foo's expiration time to 30 seconds
mc completion > mc.sh      # Generate bash completion
```

## Usage

```
$ ./mc
mc is a command line client for memcached it supports the usual
commands, such as get, set, etc.

Usage:
  mc [command]

Available Commands:
  add         Adds a key and value, if it doesn't already exist
  completion  Generates bash completion
  delete      Delete a value
  deleteall   Deletes all values
  get         Gets one of more values
  help        Help about any command
  replace     Replaces a key and value, if it already exists
  set         Sets a key and value
  touch       Touches a key (updates its expiration time)
  version     Print the version of mc

Flags:
      --config string   config file (default is $HOME/.mc.yaml)
  -h, --help            help for mc
  -p, --port string     server port (default "11211")
  -s, --server string   server hostname (default "localhost")
  -v, --verbose         verbose output

Use "mc [command] --help" for more information about a command.
```

## Configuration

`mc` can be configured with a configuration file, `$HOME/.mc.yaml`, and with
environment variables.

If multiple options are used, command line options has the highest priority,
then environment variables and the configuration file has the lowest.

### File ~/.mc.yaml

Default options `--server` and `--port` can be read from a configuration file,
default `~/.mc.yaml`.
```
# $HOME/.mc.yaml
server: memcached.myhost.com
port: 5000
```

### Environment variables

Environment variables are also supported: `SERVER` and `PORT`.

```
SERVER=my.computer.com PORT=5000 go run main.go get dingo
```

## Extended Usage

### Get

```
$ mc get -h
Gets one of more values.
	Keys should be comma separated without spaces. If no keys have values, get will exit with an error status. Multiple values will be returned one on each line in the order the keys were given. Missing keys will have the value key:none

Usage:
  mc get <key(s)> [flags]

Flags:
  -h, --help   help for get
```

### Set

```
$ mc set -h
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
$ mc delete -h
Delete a value from memcached by key

Usage:
  mc delete [flags]
```

### Delete All

```
$ mc deleteall -h
Deletes all values from memcached

Usage:
  mc deleteall [flags]
```

### Add

Adds a key and value if it doesn't exist

Usage:
  mc add [flags]

Flags:
  -e, --expiration int32   Expiration time for this key
  -f, --file string        Filename containing the value
  -h, --help               help for add

### Replace

Replaces a key and value if it already exist

Usage:
  mc replace [flags]

Flags:
  -e, --expiration int32   Expiration time for this key
  -f, --file string        Filename containing the value
  -h, --help               help for replace

### Touch

Touches a key (updates it expiration time).

Usage:
  mc touch key [flags]

Flags:
  -e, --expiration int32   Expiration time for this key
  -h, --help               help for touch

### Completion

```
$ mc completion --help
Generates bash completion, redirect to a file and move it to your
bash completion directory.

# Example
$ mc completion > mc_completion.sh # Copy to bash completion directory
```


## Development

### Install

```
go get github.com/andersjanmyr/mc
```

### Build

```
# Builds a binary for all platforms
make build
```

### Release

Builds all platforms, updates the version in `pkg/mc/version.go`, tags the
version in git, and uploads the binaries to Github releases. The changes are
not pushed to the remote.

```
# Ensure that you have updated the RELEASE_NOTES.md with information about the
release
# The release will fail if this is not included.
make release VERSION=v.1.3.2
```
