package asserter

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	assert := New(t)
	assert(true).Fail()

	assert = New(&noopT{})
	assert(false).Helper()
	assert(false).Error()
	assert(false).Errorf("%s", "yes")
	assert(false).Fatal()
	assert(false).Fatalf("%s", "yes")
	assert(false).Fail()
	assert(false).FailNow()
	assert(false).Equals(true, false)
	assert(false).Equals(true, false, "case 1")
	assert(true, false) // More than one is disallowed
}

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")
