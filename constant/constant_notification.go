package constant

type Notification int

const (
	// SysRelated.
	NotificationBegin Notification = 1000

	FriendApplicationApprovedNotification Notification = 1201 // add_friend_response
	FriendApplicationRejectedNotification Notification = 1202 // add_friend_response
	FriendApplicationNotification         Notification = 1203 // add_friend
	FriendAddedNotification               Notification = 1204
	FriendDeletedNotification             Notification = 1205 // delete_friend
	FriendRemarkSetNotification           Notification = 1206 // set_friend_remark?
	BlackAddedNotification                Notification = 1207 // add_black
	BlackDeletedNotification              Notification = 1208 // remove_black
	FriendInfoUpdatedNotification         Notification = 1209
	FriendsInfoUpdateNotification         Notification = 1210 //update friend info

	ConversationChangeNotification Notification = 1300 // change conversation opt

	UserNotificationBegin         Notification = 1301
	UserInfoUpdatedNotification   Notification = 1303 // SetSelfInfoTip              = 204
	UserStatusChangeNotification  Notification = 1304
	UserCommandAddNotification    Notification = 1305
	UserCommandDeleteNotification Notification = 1306
	UserCommandUpdateNotification Notification = 1307

	UserSubscribeOnlineStatusNotification Notification = 1308

	UserNotificationEnd Notification = 1399
	OANotification      Notification = 1400

	GroupNotificationBegin Notification = 1500

	GroupCreatedNotification                 Notification = 1501
	GroupInfoSetNotification                 Notification = 1502
	JoinGroupApplicationNotification         Notification = 1503
	MemberQuitNotification                   Notification = 1504
	GroupApplicationAcceptedNotification     Notification = 1505
	GroupApplicationRejectedNotification     Notification = 1506
	GroupOwnerTransferredNotification        Notification = 1507
	MemberKickedNotification                 Notification = 1508
	MemberInvitedNotification                Notification = 1509
	MemberEnterNotification                  Notification = 1510
	GroupDismissedNotification               Notification = 1511
	GroupMemberMutedNotification             Notification = 1512
	GroupMemberCancelMutedNotification       Notification = 1513
	GroupMutedNotification                   Notification = 1514
	GroupCancelMutedNotification             Notification = 1515
	GroupMemberInfoSetNotification           Notification = 1516
	GroupMemberSetToAdminNotification        Notification = 1517
	GroupMemberSetToOrdinaryUserNotification Notification = 1518
	GroupInfoSetAnnouncementNotification     Notification = 1519
	GroupInfoSetNameNotification             Notification = 1520

	//SignalingNotificationBegin Notification = 1600
	//SignalingNotification     Notification = 1601
	//SignalingNotificationEnd  Notification = 1649

	SuperGroupNotificationBegin  Notification = 1650
	SuperGroupUpdateNotification Notification = 1651
	MsgDeleteNotification        Notification = 1652
	SuperGroupNotificationEnd    Notification = 1699

	ConversationPrivateChatNotification Notification = 1701
	ConversationUnreadNotification      Notification = 1702
	ClearConversationNotification       Notification = 1703

	BusinessNotificationBegin Notification = 2000
	BusinessNotification      Notification = 2001
	BusinessNotificationEnd   Notification = 2099

	MsgRevokeNotification  Notification = 2101
	DeleteMsgsNotification Notification = 2102

	HasReadReceipt Notification = 2200

	NotificationEnd Notification = 5000
)
