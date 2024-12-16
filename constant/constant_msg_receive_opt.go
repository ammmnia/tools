package constant

type MsgReceiveOpt int

// MsgReceiveOpt.
const (
	ReceiveMessage          MsgReceiveOpt = 0
	NotReceiveMessage       MsgReceiveOpt = 1
	ReceiveNotNotifyMessage MsgReceiveOpt = 2
)
