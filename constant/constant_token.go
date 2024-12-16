package constant

type TokenType int

// token.
const (
	NormalToken  TokenType = 0
	InValidToken TokenType = 1
	KickedToken  TokenType = 2
	ExpiredToken TokenType = 3
)
