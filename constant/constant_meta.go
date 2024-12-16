package constant

type MetaKey string

const (
	OperationID     MetaKey = "operationID"
	OpUserID        MetaKey = "opUserID"
	ConnID          MetaKey = "connID"
	OpUserPlatform  MetaKey = "platform"
	Token           MetaKey = "token"
	RpcCustomHeader MetaKey = "customHeader" // rpc中间件自定义ctx参数
	CheckKey        MetaKey = "CheckKey"
	TriggerID       MetaKey = "triggerID"
	RemoteAddr      MetaKey = "remoteAddr"
	Account         MetaKey = "account"
)
