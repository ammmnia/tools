package constant

type GroupStatus int
type GroupType int
type UserJoinGroupSource int

const (
	// GroupStatus.
	GroupOk              GroupStatus = 0
	GroupBanChat         GroupStatus = 1
	GroupStatusDismissed GroupStatus = 2
	GroupStatusMuted     GroupStatus = 3

	// GroupType.
	NormalGroup         GroupType = 0
	SuperGroup          GroupType = 1
	WorkingGroup        GroupType = 2
	GroupBaned          GroupType = 3
	GroupBanPrivateChat GroupType = 4

	// UserJoinGroupSource.
	JoinByAdmin      UserJoinGroupSource = 1
	JoinByInvitation UserJoinGroupSource = 2
	JoinBySearch     UserJoinGroupSource = 3
	JoinByQRCode     UserJoinGroupSource = 4
)

type EnterGroupVerificationType int

const (
	ApplyNeedVerificationInviteDirectly EnterGroupVerificationType = 0 // 申请需要同意 邀请直接进
	AllNeedVerification                 EnterGroupVerificationType = 1 // 所有人进群需要验证，除了群主管理员邀请进群
	Directly                            EnterGroupVerificationType = 2 // 直接进群
)

func GroupIsBanChat(status GroupStatus) bool {
	if status != GroupStatusMuted {
		return false
	}
	return true
}

func GroupIsBanPrivateChat(gt GroupType) bool {
	if gt != GroupBanPrivateChat {
		return false
	}
	return true
}
