package constant

type ContentType int

const (
	///ContentType
	//UserRelated.
	ContentTypeBegin ContentType = 100
	Text             ContentType = 101
	Picture          ContentType = 102
	Voice            ContentType = 103
	Video            ContentType = 104
	File             ContentType = 105
	AtText           ContentType = 106
	Merger           ContentType = 107
	Card             ContentType = 108
	Location         ContentType = 109
	Custom           ContentType = 110
	Revoke           ContentType = 111
	Typing           ContentType = 113
	Quote            ContentType = 114

	AdvancedText ContentType = 117

	CustomNotTriggerConversation ContentType = 119
	CustomOnlineOnly             ContentType = 120
	ReactionMessageModifier      ContentType = 121
	ReactionMessageDeleter       ContentType = 122

	Common             ContentType = 200
	GroupMsg           ContentType = 201
	SignalMsg          ContentType = 202
	CustomNotification ContentType = 203
)

var ContentType2PushContent = map[ContentType]string{
	Picture:   "[PICTURE]",
	Voice:     "[VOICE]",
	Video:     "[VIDEO]",
	File:      "[File]",
	Text:      "[TEXT]",
	AtText:    "[@TEXT]",
	GroupMsg:  "[GROUPMSG]]",
	Common:    "[NEWMSG]",
	SignalMsg: "[SIGNALINVITE]",
}
