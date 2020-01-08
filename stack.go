package errors

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
)

const (
	stackMaxSize = 64
	stackSkip    = 1
)

type stack []uintptr

func (m stack) write(header string, w io.Writer) {

	for _, ptr := range m {
		fn := runtime.FuncForPC(ptr)
		file, line := fn.FileLine(ptr)

		_, _ = fmt.Fprintf(w, "%s%s -> %s:%d\n",
			header,
			fn.Name(),
			file,
			line,
		)
	}
}

func (m stack) String() string {
	var (
		w = bytes.NewBuffer(nil)
	)

	m.write("", w)

	return w.String()
}

func newStack(skip int) stack {

	var (
		stack [stackMaxSize]uintptr
	)

	n := runtime.Callers(
		skip,
		stack[:])

	return stack[:n]
}
