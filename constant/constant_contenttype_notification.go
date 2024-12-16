package constant

const (
	// SysRelated.
	NotificationBegin ContentType = 1000

	FriendApplicationApprovedNotification ContentType = 1201 // add_friend_response
	FriendApplicationRejectedNotification ContentType = 1202 // add_friend_response
	FriendApplicationNotification         ContentType = 1203 // add_friend
	FriendAddedNotification               ContentType = 1204
	FriendDeletedNotification             ContentType = 1205 // delete_friend
	FriendRemarkSetNotification           ContentType = 1206 // set_friend_remark?
	BlackAddedNotification                ContentType = 1207 // add_black
	BlackDeletedNotification              ContentType = 1208 // remove_black
	FriendInfoUpdatedNotification         ContentType = 1209
	FriendsInfoUpdateNotification         ContentType = 1210 //update friend info

	ConversationChangeNotification ContentType = 1300 // change conversation opt

	UserNotificationBegin         ContentType = 1301
	UserInfoUpdatedNotification   ContentType = 1303 // SetSelfInfoTip              = 204
	UserStatusChangeNotification  ContentType = 1304
	UserCommandAddNotification    ContentType = 1305
	UserCommandDeleteNotification ContentType = 1306
	UserCommandUpdateNotification ContentType = 1307

	UserSubscribeOnlineStatusNotification ContentType = 1308

	UserNotificationEnd ContentType = 1399
	OANotification      ContentType = 1400

	GroupNotificationBegin ContentType = 1500

	GroupCreatedNotification                 ContentType = 1501
	GroupInfoSetNotification                 ContentType = 1502
	JoinGroupApplicationNotification         ContentType = 1503
	MemberQuitNotification                   ContentType = 1504
	GroupApplicationAcceptedNotification     ContentType = 1505
	GroupApplicationRejectedNotification     ContentType = 1506
	GroupOwnerTransferredNotification        ContentType = 1507
	MemberKickedNotification                 ContentType = 1508
	MemberInvitedNotification                ContentType = 1509
	MemberEnterNotification                  ContentType = 1510
	GroupDismissedNotification               ContentType = 1511
	GroupMemberMutedNotification             ContentType = 1512
	GroupMemberCancelMutedNotification       ContentType = 1513
	GroupMutedNotification                   ContentType = 1514
	GroupCancelMutedNotification             ContentType = 1515
	GroupMemberInfoSetNotification           ContentType = 1516
	GroupMemberSetToAdminNotification        ContentType = 1517
	GroupMemberSetToOrdinaryUserNotification ContentType = 1518
	GroupInfoSetAnnouncementNotification     ContentType = 1519
	GroupInfoSetNameNotification             ContentType = 1520

	//SignalingNotificationBegin Notification = 1600
	//SignalingNotification     Notification = 1601
	//SignalingNotificationEnd  Notification = 1649

	SuperGroupNotificationBegin  ContentType = 1650
	SuperGroupUpdateNotification ContentType = 1651
	MsgDeleteNotification        ContentType = 1652
	SuperGroupNotificationEnd    ContentType = 1699

	ConversationPrivateChatNotification ContentType = 1701
	ConversationUnreadNotification      ContentType = 1702
	ClearConversationNotification       ContentType = 1703

	BusinessNotificationBegin ContentType = 2000
	BusinessNotification      ContentType = 2001
	BusinessNotificationEnd   ContentType = 2099

	MsgRevokeNotification  ContentType = 2101
	DeleteMsgsNotification ContentType = 2102

	HasReadReceipt ContentType = 2200

	NotificationEnd ContentType = 5000
)
