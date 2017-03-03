# epoch-test: Demo Golang oddities in handling Unix epoch timestamps

If the local timezone is not UTC, then Go's Unix epochs may format and parse incorrectly.

# EXAMPLE

```console
$ date "+%Z %z"
CST -0600

$ go test
--- FAIL: TestEpochConversion (0.00s)
	epoch_test.go:21: Expected 2016-01-01 01:00:00 +0000 UTC, got 2015-12-31 19:00:00 -0600 CST
FAIL
exit status 1
FAIL  github.com/mcandre/epoch-derp 0.007s
```
