package qual

import (
	"net/http"
	"os"
)

// When using the Assert func even though you do not have access to the
// source code, eg. when doing integration tests a generic output will be given.
func ExampleAssert_missingSourceCode() {
	os.Rename("assert_example_test.go", "missing.go")
	defer os.Rename("missing.go", "assert_example_test.go")

	msg := "Verify response status code"
	resp, _ := http.Get("http://www.example.com/existing_file")
	Assert(t, Vars{resp.StatusCode, msg},
		resp.StatusCode == 200,
	)
	// Output:
	//   failed assert[0]
	//     Vars[0] = 404
	//     Vars[1] = "Verify response status code"
}
