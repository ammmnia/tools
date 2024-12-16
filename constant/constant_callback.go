package constant

type CallbackCommandVal string
type CallbackActionCode int
type CallbackHandleCode int

const CallbackCommand = "command"

const (
	// callbackCommand.
	CallbackBeforeSendSingleMsgCommand                   CallbackCommandVal = "callbackBeforeSendSingleMsgCommand"
	CallbackAfterSendSingleMsgCommand                    CallbackCommandVal = "callbackAfterSendSingleMsgCommand"
	CallbackBeforeSendGroupMsgCommand                    CallbackCommandVal = "callbackBeforeSendGroupMsgCommand"
	CallbackAfterSendGroupMsgCommand                     CallbackCommandVal = "callbackAfterSendGroupMsgCommand"
	CallbackMsgModifyCommand                             CallbackCommandVal = "callbackMsgModifyCommand"
	CallbackUserOnlineCommand                            CallbackCommandVal = "callbackUserOnlineCommand"
	CallbackUserOfflineCommand                           CallbackCommandVal = "callbackUserOfflineCommand"
	CallbackUserKickOffCommand                           CallbackCommandVal = "callbackUserKickOffCommand"
	CallbackOfflinePushCommand                           CallbackCommandVal = "callbackOfflinePushCommand"
	CallbackOnlinePushCommand                            CallbackCommandVal = "callbackOnlinePushCommand"
	CallbackSuperGroupOnlinePushCommand                  CallbackCommandVal = "callbackSuperGroupOnlinePushCommand"
	CallbackBeforeAddFriendCommand                       CallbackCommandVal = "callbackBeforeAddFriendCommand"
	CallbackBeforeUpdateUserInfoCommand                  CallbackCommandVal = "callbackBeforeUpdateUserInfoCommand"
	CallbackBeforeCreateGroupCommand                     CallbackCommandVal = "callbackBeforeCreateGroupCommand"
	CallbackBeforeMemberJoinGroupCommand                 CallbackCommandVal = "callbackBeforeMemberJoinGroupCommand"
	CallbackBeforeSetGroupMemberInfoCommand              CallbackCommandVal = "CallbackBeforeSetGroupMemberInfoCommand"
	CallbackBeforeSetMessageReactionExtensionCommand     CallbackCommandVal = "callbackBeforeSetMessageReactionExtensionCommand"
	CallbackBeforeDeleteMessageReactionExtensionsCommand CallbackCommandVal = "callbackBeforeDeleteMessageReactionExtensionsCommand"
	CallbackGetMessageListReactionExtensionsCommand      CallbackCommandVal = "callbackGetMessageListReactionExtensionsCommand"
	CallbackAddMessageListReactionExtensionsCommand      CallbackCommandVal = "callbackAddMessageListReactionExtensionsCommand"

	// callback actionCode.
	ActionAllow     CallbackActionCode = 0
	ActionForbidden CallbackActionCode = 1
	// callback callbackHandleCode.
	CallbackHandleSuccess CallbackHandleCode = 0
	CallbackHandleFailed  CallbackHandleCode = 1
)
