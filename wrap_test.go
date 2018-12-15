package qual

import (
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	If := Wrap(t)
	If(false).Log("Not ok")
	If(true).Log("Ok")
}

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")

func ExampleWrap() {
	If := Wrap(t)
	If(err != nil).Errorf("...")
	If(true).Log("...")
	If(false).Fail()
}
