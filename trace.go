package errors

import (
	"bytes"
	"fmt"
)

type traceError struct {
	err   error
	stack stack
}

func (m *traceError) Unwrap() error {
	return m.err
}

func (m *traceError) Error() string {
	var (
		w = bytes.NewBuffer(nil)
	)

	_, _ = fmt.Fprintln(w, m.err)

	m.stack.write("    ", w)

	return w.String()
}
