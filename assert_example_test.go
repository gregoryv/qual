package qual

import (
	"os"
)

func ExampleAssert_missing_source_code() {
	os.Rename("assert_example_test.go", "missing.go")
	defer os.Rename("missing.go", "assert_example_test.go")

	msg := "Eg. integration test"
	result := make([]int, 0)
	Assert(t, Vars{len(result), msg},
		len(result) == 1,
	)
	// Output:
	//   failed assert[0]
	//     Vars[0] = 0
	//     Vars[1] = "Eg. integration test"
}
