package errs

import "testing"

func TestWrap(t *testing.T) {
	err := ErrRefuseFriend.WrapLocal(ErrRefuseFriend.Msg(), ErrRefuseFriend.Code())
	t.Log(err)
}
