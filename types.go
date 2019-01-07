package qual

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fail()
	FailNow()
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
