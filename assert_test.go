package qual

import (
	"fmt"
	"testing"
)

type mockT struct{}

func (m *mockT) Helper() {}
func (m *mockT) Error(args ...interface{}) {
	fmt.Println(args...)
}
func (m *mockT) Log(args ...interface{}) {
	fmt.Println(args...)
}

func (m *mockT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}

func TestAssertAbove(t *testing.T) {
	val, err := 1, fmt.Errorf("some error")
	Assert(&mockT{}, Vars{val, err},
		val == 1,
		err != nil,
	)
}

func Test_above(t *testing.T) {
	val := 2
	str := above(1)
	Assert(t, Vars{val, str},
		str == "val := 2",
		val == 2,
	)
}

var t = &mockT{}

func ExampleAssert() {
	// Some test expression. The line just above Assert(...) will be
	// printed for context if en error occurs.
	val, err := 1, fmt.Errorf("This is an error")
	Assert(t, Vars{val, err},
		val == 2, // each of these must be on a new line
		err != nil,
	)
	//output:
	//> val, err := 1, fmt.Errorf("This is an error")
	//failed assert: val == 2
	//> val = 1
	//> err = "This is an error"
}
