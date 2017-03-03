# epoch-test: Demo Golang oddities in handling Unix epoch timestamps

Specifically, while `time.Unix()` parses Unix epoch integers from timezone UTC,
the return time object is expressed in terms of the computer's local timezone,
which may not always be UTC.

Accordingly, users should explicitly `(*Time).In(time.UTC)` any parsed epochs, for safety.

# EXAMPLE

```console
$ date "+%Z %z"
CST -0600

$ go test
PASS
ok	github.com/mcandre/epoch-derp	0.006s
```
