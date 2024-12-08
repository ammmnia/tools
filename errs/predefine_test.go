package errs

import "testing"

func TestWrap(t *testing.T) {
	err := ErrRefuseFriend.WrapMsgKV("State", 1)
	t.Log(err)
}
