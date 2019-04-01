package qual

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
}

type noopT struct{}

func (t *noopT) Helper()                       {}
func (t *noopT) Error(...interface{})          {}
func (t *noopT) Errorf(string, ...interface{}) {}
