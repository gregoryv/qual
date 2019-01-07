package qual

import (
	"fmt"
	"testing"
)

func TestAssert(t *testing.T) {
	assert := Assert(t)
	assert(true).Log("Not ok")
	assert(false).Logf("%v", false)

	assert = Assert(&noopT{})
	assert(false).Error()
	assert(false).Errorf("%s", "yes")
	assert(false).Fatal()
	assert(false).Fatalf("%s", "yes")
	assert(false).Log()
	assert(false).Logf("%s %s", "yes", "no")
	assert(false).Fail()
	assert(false).FailNow()
	assert(false).Skip()
	assert(false).SkipNow()
	assert(false).Skipf("%s %s", "yes", "no")
}

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")

func ExampleAssert() {
	assert := Assert(t)
	assert(err != nil).Errorf("...")
	assert(false).Log("...")
	assert(true).Fail()
}
