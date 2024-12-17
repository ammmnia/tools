package constant

type OptionsKey = string

// OptionsKey.
const (
	IsHistory                  OptionsKey = "history"
	IsPersistent               OptionsKey = "persistent"
	IsOfflinePush              OptionsKey = "offlinePush"
	IsUnreadCount              OptionsKey = "unreadCount"
	IsConversationUpdate       OptionsKey = "conversationUpdate"
	IsSenderSync               OptionsKey = "senderSync"
	IsNotPrivate               OptionsKey = "notPrivate"
	IsSenderConversationUpdate OptionsKey = "senderConversationUpdate"
	IsSenderNotificationPush   OptionsKey = "senderNotificationPush"
	IsReactionFromCache        OptionsKey = "reactionFromCache"
	IsNotNotification          OptionsKey = "isNotNotification"
	IsSendMsg                  OptionsKey = "isSendMsg"
)
