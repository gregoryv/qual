/* package defines AssertFunc wrapper of testing.T

Online assertions are done by wrapping the T in a test

    func TestSomething(t *testing.T) {
        assert := asserter.New(t)
        got, err := Something()
        t.Logf("%v, %v := Something()", got, err)
        assert(err == nil).Fail()
        // Special case used very often is check equality
        assert().Equals(got, 1)
    }
*/
package asserter

import "strings"

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fail()
	FailNow()
}

type A interface {
	T
	Equals(got, exp interface{}, msg ...string)
}

type AssertFunc func(expr ...bool) A

type ta struct {
	t T
}

func (t *ta) Helper() {
	/* Cannot use the asserter as helper */
}
func (t *ta) Error(args ...interface{})                 { t.t.Helper(); t.t.Error(args...) }
func (t *ta) Errorf(format string, args ...interface{}) { t.t.Helper(); t.t.Errorf(format, args...) }
func (t *ta) Fatal(args ...interface{})                 { t.t.Helper(); t.t.Fatal(args...) }
func (t *ta) Fatalf(format string, args ...interface{}) { t.t.Helper(); t.t.Fatalf(format, args...) }
func (t *ta) Fail()                                     { t.t.Helper(); t.t.Fail() }
func (t *ta) FailNow()                                  { t.t.Helper(); t.t.FailNow() }
func (t *ta) Equals(got, exp interface{}, msg ...string) {
	t.t.Helper()
	if got != exp {
		str := ""
		if len(msg) > 0 {
			str = " " + strings.Join(msg, " ")
		}
		t.Errorf("got %v, expected %v%s", got, exp, str)
	}
}

type noopT struct{}

func (t *noopT) Helper()                                    {}
func (t *noopT) Error(...interface{})                       {}
func (t *noopT) Errorf(string, ...interface{})              {}
func (t *noopT) Fatal(...interface{})                       {}
func (t *noopT) Fatalf(string, ...interface{})              {}
func (t *noopT) Fail()                                      {}
func (t *noopT) FailNow()                                   {}
func (t *noopT) Equals(got, exp interface{}, msg ...string) {}

var ok *noopT = &noopT{}

// Assert returns an asserter for online assertions.
func New(t T) AssertFunc {
	return func(expr ...bool) A {
		if len(expr) > 1 {
			t.Helper()
			t.Fatal("Only 0 or 1 bool expressions are allowed")
		}
		if len(expr) == 0 || !expr[0] {
			return &ta{t}
		}
		return ok
	}
}
