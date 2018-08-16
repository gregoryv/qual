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
func (m *mockT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}

func TestAssertAbove(t *testing.T) {
	val, err := 1, fmt.Errorf("some error")
	AssertAbove(&mockT{},
		val == 1,
		err != nil,
	)
}

func Test_me(t *testing.T) {
	year := 2018
	var argName string
	func(x int) {
		argName = me(0)
	}(year)
	Assert(t, "me(1)",
		argName == "year",
	)
}

func TestAssert(t *testing.T) {
	val := 1
	Assert(t, "",
		val == 1,
	)
	err := fmt.Errorf("some error")
	Assert(&mockT{}, "",
		err == nil,
	)
	Assert(&mockT{}, "x",
		err != nil,
	)
}

func Test_above(t *testing.T) {
	val := 2
	str := above(1)
	AssertAbove(t,
		str == "val := 2",
		val == 2,
	)
}

var t = &mockT{}

func ExampleAssert() {
	// Some test expression
	val, err := 1, fmt.Errorf("This is an error")
	Assert(t, "Should not fail",
		val == 2,
		err == nil,
	)
	//output:
	//Should not fail
	//val == 2 false
	//err == nil false
}

func ExampleAssertAbove() {
	// Some test expression
	val, err := 1, fmt.Errorf("This is an error")
	AssertAbove(t,
		val == 2, // each of these must be on a new line
		err != nil,
	)
	//output:
	//val, err := 1, fmt.Errorf("This is an error")
	//val == 2 false
}
