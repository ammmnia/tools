package constant

type AppMangerLevel int

const (
	IMOrdinaryUser       AppMangerLevel = 0
	AppOrdinaryUsers     AppMangerLevel = 1
	AppAdmin             AppMangerLevel = 2
	AppNotificationAdmin AppMangerLevel = 3
)
