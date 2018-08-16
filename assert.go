package qual

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Vars []interface{}

func AssertAbove(t T, v Vars, oks ...bool) (failed bool) {
	t.Helper()
	return assert(t, above(2), v, oks...)
}

func assert(t T, msg string, v Vars, oks ...bool) (failed bool) {
	t.Helper()
	for i, ok := range oks {
		if !ok {
			if !failed {
				t.Errorf("%s", msg)
				failed = true
			}
			t.Errorf("assert: %s", trueCase(i+1))
		}
	}
	if failed {
		logVars(t, v, strings.Join(funcArgs(3), ","))
	}
	return
}

func logVars(t T, v Vars, parts string) {
	t.Helper()
	i := strings.Index(parts, "{") + 1
	j := strings.Index(parts, "}")
	vars := strings.Split(parts[i:j], ",")
	for i, v := range v {
		var val string
		switch v := v.(type) {
		case string, error:
			val = fmt.Sprintf("%q", v)
		default:
			val = fmt.Sprintf("%v", v)
		}
		t.Log(strings.TrimSpace(vars[i]), "=", val)
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

func funcArgs(n int) []string {
	_, file, line, _ := runtime.Caller(n) // cannot fail in this context
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
	if j == -1 {
		return strings.Split(str[i:], ",")
	}
	return strings.Split(str[i:j], ",")
}

// returns the variable name of the calling func
func me(nth int) string {
	parts := funcArgs(3)
	return strings.TrimSpace(parts[nth])
}

// returns the variable name of the calling func
func trueCase(nth int) string {
	_, file, line, _ := runtime.Caller(3) // cannot fail in this context
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
