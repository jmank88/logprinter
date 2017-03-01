package logprinter

import (
	"fmt"

	"github.com/go-kit/kit/log"
)

func ExamplePrinter() {
	var l log.Logger = &exampleLogger{}
	p := NewPrinter(l, "msg")
	p.Print("test")
	p.Printf("%d", 10)
	// Output:
	// msg test
	// msg 10
}

func ExamplePrinter_context() {
	c := log.NewContext(&exampleLogger{}).With("key","val")
	NewPrinter(c, "msg").Print("test")
	c = c.WithPrefix("prefixKey", "prefixVal")
	NewPrinter(c, "msg").Printf("%d", 10)
	// Output:
	// key val msg test
	// prefixKey prefixVal key val msg 10
}

type exampleLogger struct {}

func (l *exampleLogger) Log(keyvals ...interface{}) error {
	_, err := fmt.Println(keyvals...)
	return err
}