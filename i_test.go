package qual

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

// When using the Assert func even though you do not have access to the
// source code, eg. when doing integration tests a generic output will be given.
func TestSomeIntegration(t *testing.T) {
	os.Rename("i_test.go", "missing_i_test.go")
	defer os.Rename("missing_i_test.go", "i_test.go")

	exp := 400
	msg := fmt.Sprintf("Expect %v", exp)
	resp, _ := http.Get("http://www.example.com/")
	Assert(t, Vars{msg, resp.StatusCode},
		resp.StatusCode == exp,
	)
}
