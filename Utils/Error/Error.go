package Error

import "errors"

var (
	ErrpasswordLen           = errors.New("password error: length must be between 5 and 20 characters")
	ErrName                  = errors.New("name error: length must be between 5 and 20 characters")
	UsernameOrPassword       = errors.New("password error: password error")
	UserExists               = errors.New("UserExists: user exists")
	InternalError            = errors.New("InternalErr: internal error")
	TokenError               = errors.New("token invalid")
	RequestNull              = errors.New("request is null")
	UserNotExists            = errors.New("UserNotExists: user not exists")
	NoDocumentFoundWithToken = errors.New("NoDocumentFoundWithToken: document not found with token")
	DocumentNotFound         = errors.New("DocumentNotExists: document not exists")
	EnterForbidden           = errors.New("EnterForbidden: you dose not have permission")
)
