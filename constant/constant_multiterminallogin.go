package constant

type MultiTerminalLogin int

const (

	// MultiTerminalLogin.
	DefalutNotKick MultiTerminalLogin = 0
	// Full-end login, but the same end is mutually exclusive.
	AllLoginButSameTermKick MultiTerminalLogin = 1
	// Only one of the endpoints can log in.
	SingleTerminalLogin MultiTerminalLogin = 2
	// The web side can be online at the same time, and the other side can only log in at one end.
	WebAndOther MultiTerminalLogin = 3
	// The PC side is mutually exclusive, and the mobile side is mutually exclusive, but the web side can be online at
	// the same time.
	PcMobileAndWeb MultiTerminalLogin = 4
	// The PC terminal can be online at the same time,but other terminal only one of the endpoints can login.
	PCAndOther MultiTerminalLogin = 5
	// Customize login policy
	Customize MultiTerminalLogin = 6
)
