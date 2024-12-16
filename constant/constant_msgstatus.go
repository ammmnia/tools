package constant

type MsgSendingStatus int

const (
	MsgStatusSending     MsgSendingStatus = 1
	MsgStatusSendSuccess MsgSendingStatus = 2
	MsgStatusSendFailed  MsgSendingStatus = 3
	MsgStatusHasDeleted  MsgSendingStatus = 4
	MsgStatusFiltered    MsgSendingStatus = 5
)
