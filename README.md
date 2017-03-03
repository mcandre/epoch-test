# epoch-test: Demo Golang oddities in handling Unix epoch timestamps

If the local timezone is not UTC, then Go's Unix epochs may format and parse incorrectly.

# EXAMPLE

```console
$ date "+%Z %z"
CST -0600

$ go test
PASS
ok	github.com/mcandre/epoch-derp	0.006s
```
