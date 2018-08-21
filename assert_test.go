package qual

import (
	"fmt"
	"os"
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

func TestAssert(t *testing.T) {
	val, err := 1, fmt.Errorf("some error")
	Assert(&mockT{}, Vars{val, err},
		val == 1,
		err != nil,
	)
}

var t = &mockT{}

func ExampleAssert() {
	// In this example the Assert call fails to show the output.  The
	// line just above Assert(...) will be printed for context if an
	// error occurs.
	val, err := []byte("an apple"), fmt.Errorf("This is an error")
	Assert(t, Vars{val, err},
		len(val) == 10,
		err != nil,
	)
	// Output:
	// > val, err := []byte("an apple"), fmt.Errorf("This is an error")
	//   failed assert: len(val) == 10
	//     val = "an apple"
	//     err = "This is an error"
}

func ExampleAssert_oneline() {
	val, err := 1, fmt.Errorf("This is an error")
	Assert(t, Vars{val, err}, val == 2, err == nil)
	// Output:
	// > val, err := 1, fmt.Errorf("This is an error")
	//   failed assert: val == 2
	//   failed assert: err == nil
	//     val = 1
	//     err = "This is an error"
}

func ExampleAssert_nil() {
	var err error
	val, err := 9, nil
	Assert(t, Vars{val, err},
		val == 9,
		err != nil,
	)
	// Output:
	// > val, err := 9, nil
	//   failed assert: err != nil
	//     val = 9
	//     err = nil
}

func Test_scanLine(t *testing.T) {
	str, err := scanLine(29, 0)
	Assert(t, Vars{str, err},
		str == "",
		err != nil,
	)

	msg := "File does not exist"
	os.Rename("assert_test.go", "m.go")
	defer os.Rename("m.go", "assert_test.go")
	str, err = scanLine(1, 0)
	Assert(t, Vars{msg, str, err},
		str == "",
		err != nil,
	)
}
