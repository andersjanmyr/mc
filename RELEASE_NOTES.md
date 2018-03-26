# mc, Release Notes

## Release v1.4.1

* Use get for single key and get many for multiple keys

## Release v1.4.0

* Added support reading value from stdin for `set`, `add`, and `replace`.
* `echo -n value | mc set key`

## Release v1.3.3

* Clarified help texts.

## Release v1.3.2

* Added support for touching a key with `mc touch key1 -e 20`.

## Release v1.3.1

* Added support for getting multiple keys with `mc get key1,key2`.

## Release v1.3.0

* Added `mc add` and `mc replace` commands.
*
## Release v1.2.2

* Added verbose option (`-v` or `--verbose`)

## Release v1.2.1

* Better error reporting for errors in config file, `$HOME/.mc.yaml`.

## Release v1.2.0

* Added `mc completion` command to generate bash completion.

## Release v1.1.0

* Renamed option host to --server (-s) to avoid conflict with --help (-h).

## Release v1.0.0

* Initial release with support for `get`, `set`, `delete`, `deleteall` and
  `version`

