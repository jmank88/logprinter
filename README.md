# logprinter [![GoDoc](https://godoc.org/github.com/jmank88/logprinter?status.svg)](https://godoc.org/github.com/jmank88/logprinter)
A Print/Printf adapter for kit/log

```go
var l log.Logger = &exampleLogger{}
p := logprinter.NewPrinter(l, "msg")
p.Print("test")
p.Printf("%d", 10)
```

Output:

```text
msg test
msg 10
```