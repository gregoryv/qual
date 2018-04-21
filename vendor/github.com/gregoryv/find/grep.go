package find

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

// Ref points to a line in a file starting from 0
type Ref struct {
	LineNo int
	Line   string
}

func (ref *Ref) String() string {
	return fmt.Sprintf("%v:%s", ref.LineNo, ref.Line)
}

// InFile opens and the file and uses InStream to find references
func InFile(pattern, file string) (result *list.List, err error) {
	var stream *os.File
	stream, err = os.Open(file)
	if err != nil {
		return
	}
	defer stream.Close()
	return InStream(pattern, stream), nil
}

// InStream works much like grep, finding references by the given pattern
func InStream(pattern string, stream io.Reader) *list.List {
	var (
		scanner = bufio.NewScanner(stream)
		sp      = NewShellPattern("*" + pattern + "*")
		line    string
		lineNo  int
		result  = list.New()
	)

	for scanner.Scan() {
		line = scanner.Text()
		if sp.Match(line) {
			result.PushBack(&Ref{lineNo, line})
		}
		lineNo++
	}
	return result
}
