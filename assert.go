package qual

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Vars groups variables which will be logged on test failures
type Vars []interface{}

func Assert(t T, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	str, err := scanLine(2, -1)
	str = strings.TrimSpace(str)
	generic := err != nil
	return assert(t, str, generic, v, checks...)
}

func assert(t T, msg string, generic bool, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	for i, ok := range checks {
		if ok {
			continue
		}
		if !failed {
			if !generic {
				t.Errorf("> %s", msg)
			}
			failed = true
		}
		if generic {
			t.Errorf("  failed assert[%v]", i)
		} else {
			str, _ := scanLine(3, 0)
			t.Errorf("  failed assert: %s", trueCase(str, i+1))
		}
	}
	if failed {
		// Log all Vars{...} with name and value
		var args []string
		var varsLine string
		if !generic {
			args, _ = funcArgs(3)
			varsLine = strings.Join(args, ",")
		} else {
			args = make([]string, 0)
			for i, _ := range v {
				args = append(args, fmt.Sprintf("Vars[%v]", i))
			}
			varsLine = fmt.Sprintf("Vars{%s}", strings.Join(args, ","))
		}
		logVars(t, v, varsLine)
	}
	return
}

// returns the variable name of the calling func
func trueCase(str string, nth int) string {
	if assertStatementIsOnSameLine(str) {
		i := strings.Index(str, "}")
		j := strings.Index(str, ")")
		parts := strings.Split(str[i:j], ",")
		return strings.TrimSpace(parts[nth])
	}
	str, _ = scanLine(4, nth) // todo handle error
	// Assuming they are on the same line here
	j := strings.Index(str, ",")
	// if j is -1 then the compiler should fail
	return strings.TrimSpace(str[:j])
}

func assertStatementIsOnSameLine(str string) bool {
	return strings.Index(str, "Assert(") >= 0 && strings.LastIndex(str, ")") > 0
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
		if v == nil {
			val = "nil"
		}
		t.Log("   ", strings.TrimSpace(vars[i]), "=", val)
	}
}

func funcArgs(n int) (args []string, err error) {
	str, err := scanLine(n+1, 0) // todo handle error
	if err != nil {
		return
	}
	i := strings.Index(str, "(") + 1
	// Assuming they are on the same line here
	j := strings.Index(str, ")")
	if j == -1 {
		args = strings.Split(str[i:], ",")
	} else {
		args = strings.Split(str[i:j], ",")
	}
	return
}

// scanLine returns the line above the caller if back == 0. Fails if source
// is not available.
func scanLine(caller, back int) (string, error) {
	_, file, line, ok := runtime.Caller(caller)
	if !ok {
		return "", fmt.Errorf("Unknown caller")
	}
	fh, err := os.Open(file)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line+back; i++ {
		scanner.Scan()
	}
	fh.Close()
	return scanner.Text(), nil
}
