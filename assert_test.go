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
	AssertAbove(&mockT{}, Vars{val, err},
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
	AssertAbove(t, Vars{argName},
		argName == "year",
	)
}
func Test_above(t *testing.T) {
	val := 2
	str := above(1)
	AssertAbove(t, Vars{val, str},
		str == "val := 2",
		val == 2,
	)
}

var t = &mockT{}

func ExampleAssertAbove() {
	// Some test expression
	val, err := 1, fmt.Errorf("This is an error")
	AssertAbove(t, Vars{val, err},
		val == 2, // each of these must be on a new line
		err != nil,
	)
	//output:
	//val, err := 1, fmt.Errorf("This is an error")
	//assert: val == 2
	//val = 1
	//err = "This is an error"
}
