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

type exampleLogger struct {}

func (l *exampleLogger) Log(keyvals ...interface{}) error {
	_, err := fmt.Println(keyvals...)
	return err
}