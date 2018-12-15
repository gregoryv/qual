package qual

import "fmt"

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")

func ExampleWrap() {
	If := Wrap(t)
	If(err != nil).Errorf("...")
	If(true).Log("...")
	If(false).Fail()
}
