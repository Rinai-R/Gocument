package ErrCode

var (
	// OK 成功
	OK = 20000
	// BindError 数据绑定出现错误
	BindError = 40000
	// UserNameLengthError 用户名长度不对
	UserNameLengthError = 40001
	// PasswordLengthError 密码长度错误
	PasswordLengthError = 40002
	// UserNameExists 该用户名已经存在
	UserNameExists = 40003
	// UsernameOrPassword 用户名或者密码错误
	UsernameOrPassword = 40004
	// TokenError token异常
	TokenError = 40005
	// RequestNull 请求数据为空
	RequestNull = 40006
	// UserNotExists 该用户不存在
	UserNotExists = 40007
	// NoDocumentFoundWithToken 根据token没有找到当前文档，文档与账号不符合
	NoDocumentFoundWithToken = 40008
	// DocumentNotFound 没有找到文档
	DocumentNotFound = 40009
	// EnterForbidden 禁止访问文档
	EnterForbidden = 40010
	// GrantFailed 授权失败
	GrantFailed = 40011
	// SensitiveWords 敏感词审查
	SensitiveWords = 40012
	// InternalErr 内部其他错误
	InternalErr = 50000
)
