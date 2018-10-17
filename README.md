# yj

The `yj` command reads YAML from stdin (or a file) and writes JSON to stdout (or to a file with the `-o` flag).

Convert JSON to YAML with [jy](https://github.com/sourcegraph/jy).

## Examples

```sh
$ echo "hello: world" | yj
{"hello":"world"}

$ yj sample.yml -o result.json
```

## Install

Install the latest with `go get -u github.com/sourcegraph/yj`.
