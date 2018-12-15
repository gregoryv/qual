package qual

import (
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	If := Wrap(t)
	If(false).Log("Not ok")
	If(true).Logf("%v", true)

	If = Wrap(&noopT{})
	If(true).Error()
	If(true).Errorf("%s", "yes")
	If(true).Fatal()
	If(true).Fatalf("%s", "yes")
	If(true).Log()
	If(true).Logf("%s %s", "yes", "no")
	If(true).Fail()
	If(true).FailNow()
	If(true).Skip()
	If(true).SkipNow()
	If(true).Skipf("%s %s", "yes", "no")
}

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")

func ExampleWrap() {
	If := Wrap(t)
	If(err != nil).Errorf("...")
	If(true).Log("...")
	If(false).Fail()
}
