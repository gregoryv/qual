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
	if err != nil {
		// If a scanning error occured use a generic error message
		return assertNoScan(t, str, v, checks...)
	}
	return assert(t, str, v, checks...)
}

func assertNoScan(t T, msg string, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	for i, ok := range checks {
		if ok {
			continue
		}
		if !failed {
			failed = true
		}
		t.Errorf("  failed assert[%v]", i)
	}
	if failed {
		varsLine := generateVarsLine(len(v))
		logVars(t, v, varsLine)
	}
	return
}

func generateVarsLine(n int) string {
	var args []string
	args = make([]string, 0)
	for i := 0; i < n; i++ {
		args = append(args, fmt.Sprintf("Vars[%v]", i))
	}
	return fmt.Sprintf("Vars{%s}", strings.Join(args, ","))
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
		str, _ := scanLine(3, 0)
		t.Errorf("  failed assert: %s", trueCase(str, i+1))
	}
	if failed {
		// Log all Vars{...} with name and value
		args := funcArgs(3)
		varsLine := strings.Join(args, ",")
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
	str, _ = scanLine(4, nth)
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

func funcArgs(n int) (args []string) {
	str, _ := scanLine(n+1, 0)
	// Scanning failures should be caught in Assert so we don't handle it here

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
