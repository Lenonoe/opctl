## `pkg install [OPTIONS] PKG_REF` (since v0.1.19)

Installs a package

## Arguments

### `PKG_REF`
Package reference (`host/path/repo#tag`)

## Options

### `--path` *default: `.opspec`*
Path to install the package at

### `-u` or `--username`
Username used to auth w/ the pkg source

### `-p` or `--password`
Password used to auth w/ the pkg source

## Examples

```shell
opctl pkg install -u someUser -p somePass host/path/repo#tag
```

## Notes

### username/password prompt

If auth failure occurs the cli will (re)prompt for username & password & re-attempt the install.

> in non-interactive terminals, the cli will note that it can't prompt due to being in a
> non-interactive terminal and exit with a non zero exit code.