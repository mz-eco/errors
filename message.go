package errors

import "fmt"

type msgError struct {
	err error
	msg string
}

func (m *msgError) Error() string {
	return fmt.Sprintf("%s\n%s", m.msg, m.err)
}

func (m *msgError) Unwrap() error {
	return m.err
}
