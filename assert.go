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
	str, _ := scanLine(2, -1)
	str = strings.TrimSpace(str)
	return assert(t, str, v, checks...)
}

func assert(t T, msg string, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	for i, ok := range checks {
		if ok {
			continue
		}
		if !failed {
			t.Errorf("> %s", msg)
			failed = true
		}
		str, _ := scanLine(3, 0) // todo handle error
		t.Errorf("  failed assert: %s", trueCase(str, i+1))
	}
	if failed {
		// Log all Vars{...} with name and value
		logVars(t, v, strings.Join(funcArgs(3), ","))
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

func funcArgs(n int) []string {
	str, _ := scanLine(n+1, 0) // todo handle error
	i := strings.Index(str, "(") + 1
	// Assuming they are on the same line here
	j := strings.Index(str, ")")
	if j == -1 {
		return strings.Split(str[i:], ",")
	}
	return strings.Split(str[i:j], ",")
}

func scanLine(caller, back int) (string, error) {
	_, file, line, ok := runtime.Caller(caller) // todo, handle error
	if !ok {
		return "", fmt.Errorf("Unknown caller")
	}
	fh, _ := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line+back; i++ {
		scanner.Scan()
	}
	fh.Close()
	return scanner.Text(), nil
}
