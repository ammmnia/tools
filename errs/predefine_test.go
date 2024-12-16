package errs

import "testing"

func TestWrap(t *testing.T) {
	err := ErrArgs.WrapLocal(ErrArgs.Msg(), "token")
	t.Log(err)
}
