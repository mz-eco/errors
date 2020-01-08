package errors

import (
	"testing"
)

func TestTrace(t *testing.T) {

	Is(New("err"), New("err"))

	t.Log(Trace(New("Trace Error")))

	t.Log(New("New Error"))

	t.Log(Wrap(New("WrapError"), "WrapFormat"))
}
