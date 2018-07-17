# yj

The `yj` command reads YAML from stdin and writes JSON to stdout (or to a file with the `-o` flag).

Convert JSON to YAML with [jy](https://github.com/sourcegraph/jy).

## Example

```sh
$ echo "hello: world" | yj
{"hello":"world"}
```

## Install

Install the latest with `go get -u github.com/sourcegraph/yj`.
