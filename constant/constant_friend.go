package constant

type BecomeFriendType int
type FriendResponse int

const (
	BecomeFriendByImport BecomeFriendType = 1 // 管理员导入
	BecomeFriendByApply  BecomeFriendType = 2 // 申请添加

	FriendResponseNotHandle FriendResponse = 0
	FriendResponseAgree     FriendResponse = 1
	FriendResponseRefuse    FriendResponse = -1
)
