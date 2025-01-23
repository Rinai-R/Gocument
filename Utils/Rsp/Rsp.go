package Rsp

import (
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
)

func Success(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.OK,
		Info: "OK",
		Data: data,
	}
}

func BindErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.BindError,
		Info: "BindError",
		Data: data,
	}
}

func UserNameLengthErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UserNameLengthError,
		Info: "UserNameLengthError",
		Data: data,
	}
}
func PasswordLengthErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.PasswordLengthError,
		Info: "PasswordLengthError",
		Data: data,
	}
}

func UsernameOrPassword(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UsernameOrPassword,
		Info: "UsernameOrPasswordError",
		Data: data,
	}
}
func UserNameExistsErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UserNameExists,
		Info: "UserNameExists",
		Data: data,
	}
}

func InternalError(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.InternalErr,
		Info: "InternalError",
		Data: data,
	}
}

func TokenError(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.TokenError,
		Info: "TokenError",
		Data: data,
	}
}

func RequestNull(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.RequestNull,
		Info: "RequestNull",
		Data: data,
	}
}

func UserNotExists(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UserNotExists,
		Info: "UserNotExists",
		Data: data,
	}
}
