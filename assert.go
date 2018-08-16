package qual

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

// Assert fails the given test if there are any non nil errors
func Assert(t T, msg string, oks ...bool) {
	t.Helper()
	var failed bool
	for i, ok := range oks {
		if !ok {
			if !failed {
				t.Errorf("%s", msg)
				failed = true
			}
			t.Errorf("%s false", trueCase(i+1))
		}
	}
}

func AssertAbove(t T, oks ...bool) {
	t.Helper()
	var failed bool
	for i, ok := range oks {
		if !ok {
			if !failed {
				t.Errorf("%s", above(2))
				failed = true
			}
			t.Errorf("%s false", trueCase(i+1))
		}
	}
}

// returns the line above the caller
func above(nth int) string {
	_, file, line, _ := runtime.Caller(nth)
	fh, _ := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line-1; i++ {
		scanner.Scan()
	}
	fh.Close()
	str := scanner.Text()
	return strings.TrimSpace(str)
}

// returns the variable name of the calling func
func me(nth int) string {
	_, file, line, _ := runtime.Caller(2) // cannot fail in this context
	fh, _ := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line; i++ {
		scanner.Scan()
	}
	fh.Close()
	str := scanner.Text()
	i := strings.Index(str, "(") + 1
	// Assuming they are on the same line here
	j := strings.Index(str, ")")
	parts := strings.Split(str[i:j], ",")
	return strings.TrimSpace(parts[nth])
}

// returns the variable name of the calling func
func trueCase(nth int) string {
	_, file, line, _ := runtime.Caller(2) // cannot fail in this context
	fh, _ := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line+nth; i++ {
		scanner.Scan()
	}
	fh.Close()
	str := scanner.Text()
	// Assuming they are on the same line here
	j := strings.Index(str, ",")
	return strings.TrimSpace(str[:j])
}
