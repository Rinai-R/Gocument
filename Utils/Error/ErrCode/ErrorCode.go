package ErrCode

var (
	OK                       = 20000
	BindError                = 40000
	UserNameLengthError      = 40001
	PasswordLengthError      = 40002
	UserNameExists           = 40003
	UsernameOrPassword       = 40004
	InternalErr              = 50000
	TokenError               = 40005
	RequestNull              = 40006
	UserNotExists            = 40007
	NoDocumentFoundWithToken = 40008
)
