package errs

import "testing"

func TestWrap(t *testing.T) {
	err := ErrPassword.WrapLocal()
	t.Log(err)
}
