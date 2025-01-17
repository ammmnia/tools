package constant

type MetaKey = string

const (
	OperationID      = "operationID"
	OpUserID         = "opUserID"
	ConnID           = "connID"
	OpUserPlatform   = "platform"
	Token            = "token"
	RpcCustomHeader  = "customHeader" // rpc中间件自定义ctx参数
	CheckKey         = "CheckKey"
	TriggerID        = "triggerID"
	RemoteAddr       = "remoteAddr"
	SuperAdminUserId = "superAdminUserId"
)
