package constant

type MsgStatus int
type MsgFrom int
type SessionType int
type SendMsgStaus int

const (
	// status.
	MsgNormal  MsgStatus = 1
	MsgDeleted MsgStatus = 4

	// MsgFrom.
	UserMsgType MsgFrom = 100
	SysMsgType  MsgFrom = 200

	// SessionType.
	SingleChatType SessionType = 1

	// sendMsgStaus.
	MsgStatusNotExist SendMsgStaus = 0
	MsgIsSending      SendMsgStaus = 1
	MsgSendSuccessed  SendMsgStaus = 2
	MsgSendFailed     SendMsgStaus = 3
)
