package Rsp

import (
	"github.com/Rinai-R/Gocument/Utils/Error"
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
		Info: Error.ErrName.Error(),
		Data: data,
	}
}
func PasswordLengthErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.PasswordLengthError,
		Info: Error.ErrpasswordLen.Error(),
		Data: data,
	}
}

func UsernameOrPassword(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UsernameOrPassword,
		Info: Error.UsernameOrPassword.Error(),
		Data: data,
	}
}
func UserNameExistsErr(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UserNameExists,
		Info: Error.UserExists.Error(),
		Data: data,
	}
}

func InternalError(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.InternalErr,
		Info: Error.InternalError.Error(),
		Data: data,
	}
}

func TokenError(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.TokenError,
		Info: Error.TokenError.Error(),
		Data: data,
	}
}

func RequestNull(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.RequestNull,
		Info: Error.RequestNull.Error(),
		Data: data,
	}
}

func UserNotExists(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.UserNotExists,
		Info: Error.UserExists.Error(),
		Data: data,
	}
}

func NoDocumentFoundWithToken(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.NoDocumentFoundWithToken,
		Info: Error.NoDocumentFoundWithToken.Error(),
		Data: data,
	}
}

func DocumentNotFound(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.DocumentNotFound,
		Info: Error.DocumentNotFound.Error(),
		Data: data,
	}
}

func EnterForbidden(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.EnterForbidden,
		Info: Error.EnterForbidden.Error(),
		Data: data,
	}
}

func GrantFailed(data interface{}) Rsp {
	return Rsp{
		Code: ErrCode.GrantFailed,
		Info: Error.GrantFailed.Error(),
		Data: data,
	}
}
