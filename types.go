package qual

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
}
