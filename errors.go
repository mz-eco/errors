package errors

import (
	"errors"
	"fmt"
)

const (
	skipLevel = 2
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	inner := errors.Unwrap(err)

	switch inner.(type) {
	case *traceError:
		return Unwrap(inner)
	case *msgError:
		return Unwrap(inner)
	default:
		return inner
	}
}

func Mark(format string, v ...interface{}) error {
	return errors.New(
		fmt.Sprintf(format, v...))
}

func Trace(err error) error {

	switch err.(type) {
	case *traceError:
		return err
	default:
		return &traceError{
			err:   err,
			stack: newStack(2),
		}
	}
}

func Wrap(err error, format string, v ...interface{}) error {

	tr, ok := err.(*traceError)

	if ok {
		return &traceError{
			err: &msgError{
				err: tr.err,
				msg: fmt.Sprintf(format, v...),
			},
			stack: tr.stack,
		}
	}

	return &traceError{
		err: &msgError{
			err: err,
			msg: fmt.Sprintf(format, v...),
		},
		stack: newStack(skipLevel),
	}
}

func New(format string, v ...interface{}) error {
	return &traceError{
		err:   fmt.Errorf(format, v...),
		stack: newStack(skipLevel),
	}
}
